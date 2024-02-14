package application

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/0xivanov/lime-ethereum-fetcher-go/db"
	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/0xivanov/lime-ethereum-fetcher-go/repo"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go"
	pc "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
)

/*
*

	Integration tests
	These tests fire up the whole app and test the entire flow

*
*/

/*
*
Test case: /ping

*
*/
func TestPingRoute(t *testing.T) {
	ctx := context.Background()
	app, postgresC, redisC := Setup(t, ctx)
	defer app.Stop()
	defer postgresC.Terminate(ctx)
	defer redisC.Terminate(ctx)
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "9090" // Default port if not provided
	}

	url := "http://localhost:" + port + "/ping"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	expected := `{"message":"pong"}`
	body := make([]byte, len(expected))
	resp.Body.Read(body)
	assert.Equal(t, expected, string(body))
}

/*
*

Test case: /authenticate -> /eth?transactionHashes -> /all

*
*/
func TestGetTransactionsFlow(t *testing.T) {
	ctx := context.Background()
	app, postgresC, redisC := Setup(t, ctx)
	defer app.Stop()
	defer postgresC.Terminate(ctx)
	defer redisC.Terminate(ctx)
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "9090" // Default port if not provided
	}

	// authenticate
	url := "http://localhost:" + port + "/lime/authenticate"
	body := strings.NewReader(`{"username":"alice","password":"alice"}`)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var response struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		t.Fatal(err)
	}
	token := response.Token

	// get transaction from ethereum and save them to db
	url = "http://localhost:" + port + "/lime/eth?transactionHashes=0x9b2f6a3c2e1aed2cccf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524"
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// check if the transaction is saved to db
	url = "http://localhost:" + port + "/lime/all"
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var transactionResponse model.TransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&transactionResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(transactionResponse.Transactions), 1)
}

func Setup(t *testing.T, ctx context.Context) (*App, tc.Container, tc.Container) {
	gin.SetMode(gin.TestMode)
	testDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current directory: %v", err)
	}

	// Get the root directory of the project
	rootDir := filepath.Dir(testDir)

	// Load environment variables from .env file in the root directory
	err = godotenv.Load(filepath.Join(rootDir, ".env"))
	if err != nil {
		t.Errorf("could not load env vars: %v", err)
	}

	port := os.Getenv("API_PORT")
	ethNodeUrl := os.Getenv("ETH_NODE_URL")
	wsEthNodeUrl := os.Getenv("WS_ETH_NODE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")

	// init logger
	l := hclog.Default()

	// init db
	pgC, pgHost, pgPort, err := startPostgreSQLContainer(t, ctx)
	if err != nil {
		t.Fatalf("failed to start PostgreSQL container: %v", err)
	}
	connStr := constructConnectionString(pgHost, pgPort)
	// time.Sleep(1 * time.Second)
	postgresDb, err := db.NewDatabse(postgres.Open(connStr))
	if err != nil {
		t.Fatalf("error connecting to postgres db")
	}

	// Start the Redis container
	redisC, redisHost, redisPort, err := startRedisContainer(t, ctx)
	if err != nil {
		t.Fatalf("failed to start Redis container: %v", err)
	}

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})

	// init repos
	transactionRepo := repo.NewTransaction(postgresDb.GetDb(), redisClient, l)
	contractRepo := repo.NewContract(postgresDb.GetDb(), l)

	// init eth clients
	client, err := ethclient.Dial(ethNodeUrl)
	wsClient, err := ethclient.Dial(wsEthNodeUrl)
	if err != nil {
		log.Fatal("error connecting to eth node")
		panic(err)
	}

	// create and start the app
	app := New(gin.Default(), port, ethNodeUrl, jwtSecret, client, wsClient, postgresDb, l, transactionRepo, contractRepo)

	go app.Start()
	return app, pgC, redisC
}

func startPostgreSQLContainer(t *testing.T, ctx context.Context) (tc.Container, string, string, error) {
	// Define the PostgreSQL container configuration
	pgC, err := pc.RunContainer(ctx,
		tc.WithImage("postgres:latest"),
		pc.WithDatabase("postgres"),
		pc.WithUsername("root"),
		pc.WithPassword("password"),
		tc.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		return nil, "", "", err
	}
	// Get the PostgreSQL container's host and port
	pgHost, err := pgC.Host(ctx)
	if err != nil {
		return nil, "", "", err
	}
	pgPort, err := pgC.MappedPort(ctx, "5432")
	if err != nil {
		return nil, "", "", err
	}

	return pgC, pgHost, pgPort.Port(), nil
}

func constructConnectionString(host, port string) string {
	return fmt.Sprintf("postgresql://root:password@%s:%s/postgres", host, port)
}

func startRedisContainer(t *testing.T, ctx context.Context) (tc.Container, string, string, error) {
	req := tc.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections").WithStartupTimeout(30 * time.Second),
	}

	redisC, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, "", "", err
	}

	// Get the Redis container's host and port
	redisHost, err := redisC.Host(ctx)
	if err != nil {
		return nil, "", "", err
	}
	redisPort, err := redisC.MappedPort(ctx, "6379")
	if err != nil {
		return nil, "", "", err
	}

	return redisC, redisHost, redisPort.Port(), nil
}
