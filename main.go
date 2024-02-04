package main

import (
	"log"
	"os"

	"github.com/0xivanov/lime-ethereum-fetcher-go/application"
	"github.com/0xivanov/lime-ethereum-fetcher-go/db"
	"github.com/0xivanov/lime-ethereum-fetcher-go/repo"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
)

func main() {
	// get env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("API_PORT")
	dbConnectionString := os.Getenv("DB_CONNECTION_URL")

	// initialize logger, db and repos
	l := hclog.Default()
	// TODO make this halt the app if the db is not present
	dbConnection, err := db.NewDatabse(postgres.Open(dbConnectionString))
	transactionRepo := repo.NewTransaction(dbConnection.GetDb(), l)

	// create and start the app
	app := application.New(gin.Default(), port, dbConnection, l, transactionRepo)
	app.Start()
}
