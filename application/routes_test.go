package application

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/0xivanov/lime-ethereum-fetcher-go/db"
	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/0xivanov/lime-ethereum-fetcher-go/repo"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	_ "github.com/proullon/ramsql/driver"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
)

/*
*

	Integration tests
	These tests fire up the whole app and test the entire flow

*
*/
func TestPingRoute(t *testing.T) {
	app, ramdb := Setup(t)
	defer app.Stop()
	defer ramdb.Close()
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

func TestGetTransactionsFlow(t *testing.T) {
	app, ramdb := Setup(t)
	defer app.Stop()
	defer ramdb.Close()
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

func Setup(t *testing.T) (*App, *sql.DB) {
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
	ramdb, err := sql.Open("ramsql", "TestGormQuickStart")
	if err != nil {
		t.Fatalf("error creating test database: %v", err)
	}
	postgres, err := db.NewDatabse(postgres.New(postgres.Config{
		Conn: ramdb,
	}))
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// init repos
	transactionRepo := repo.NewTransaction(postgres.GetDb(), redis, l)
	contractRepo := repo.NewContract(postgres.GetDb(), l)

	// init eth clients
	client, err := ethclient.Dial(ethNodeUrl)
	wsClient, err := ethclient.Dial(wsEthNodeUrl)
	if err != nil {
		log.Fatal("error connecting to eth node")
		panic(err)
	}

	// create and start the app
	app := New(gin.Default(), port, jwtSecret, ethNodeUrl, client, wsClient, postgres, l, transactionRepo, contractRepo)

	go app.Start()
	return app, ramdb
}
