package main

import (
	"log"
	"os"

	"github.com/0xivanov/lime-ethereum-fetcher-go/application"
	"github.com/0xivanov/lime-ethereum-fetcher-go/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("API_PORT")
	dbConnectionString := os.Getenv("DB_CONNECTION_URL")
	dbConnection, err := db.NewDatabse(dbConnectionString)
	// userRouter := handler.NewUser(log, validator.New())
	log := log.New(os.Stdout, "lime-ethereum-fetcher ", log.LstdFlags)
	app := application.New(gin.Default(), port, dbConnection, log)
	app.Start()
}
