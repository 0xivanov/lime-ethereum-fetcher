package main

import (
	"log"
	"os"

	"github.com/0xivanov/lime-ethereum-fetcher-go/application"
	"github.com/0xivanov/lime-ethereum-fetcher-go/db"
	"github.com/0xivanov/lime-ethereum-fetcher-go/repo"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
)

func main() {
	// get env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
		panic(err)
	}
	port := os.Getenv("API_PORT")
	postgresConnString := os.Getenv("DB_CONNECTION_URL")
	ethNodeUrl := os.Getenv("ETH_NODE_URL")
	wsEthNodeUrl := os.Getenv("WS_ETH_NODE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")

	// init logger
	l := hclog.Default()

	// init db
	postgres, err := db.NewDatabse(postgres.Open(postgresConnString))
	if err != nil {
		log.Fatalf("error connecting to postgres db: %v", err)
		panic(err)
	}
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
		log.Fatalf("error connecting to eth node %v", err)
		panic(err)
	}

	// create and start the app
	app := application.New(gin.Default(), port, ethNodeUrl, jwtSecret, client, wsClient, postgres, l, transactionRepo, contractRepo)
	app.Start()
}
