package application

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/0xivanov/lime-ethereum-fetcher-go/db"
	"github.com/0xivanov/lime-ethereum-fetcher-go/repo"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	tc "github.com/testcontainers/testcontainers-go"
	pc "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
)

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
	pgC, pgHost, pgPort, err := StartPostgreSQLContainer(t, ctx)
	if err != nil {
		t.Fatalf("failed to start PostgreSQL container: %v", err)
	}
	connStr := ConstructConnectionString(pgHost, pgPort)
	// time.Sleep(1 * time.Second)
	postgresDb, err := db.NewDatabse(postgres.Open(connStr))
	if err != nil {
		t.Fatalf("error connecting to postgres db")
	}

	// Start the Redis container
	redisC, redisHost, redisPort, err := StartRedisContainer(t, ctx)
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

func StartPostgreSQLContainer(t *testing.T, ctx context.Context) (tc.Container, string, string, error) {
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

func ConstructConnectionString(host, port string) string {
	return fmt.Sprintf("postgresql://root:password@%s:%s/postgres", host, port)
}

func StartRedisContainer(t *testing.T, ctx context.Context) (tc.Container, string, string, error) {
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
