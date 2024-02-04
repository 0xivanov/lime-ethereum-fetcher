package application

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/0xivanov/lime-ethereum-fetcher-go/db"
	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/0xivanov/lime-ethereum-fetcher-go/repo"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	_ "github.com/proullon/ramsql/driver"
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

	url := "http://localhost:" + port + "/lime/eth?transactionHashes=0x9b2f6a3c2e1aed2cccf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524"

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
	var transactions []model.Transaction
	err = json.NewDecoder(resp.Body).Decode(&transactions)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(transactions), 1)
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

	l := hclog.Default()
	ramdb, err := sql.Open("ramsql", "TestGormQuickStart")
	if err != nil {
		t.Fatalf("Error creating test database: %v", err)
	}
	dbConnection, err := db.NewDatabse(postgres.New(postgres.Config{
		Conn: ramdb,
	}))
	transactionRepo := repo.NewTransaction(dbConnection.GetDb(), l)
	app := New(gin.Default(), port, dbConnection, l, transactionRepo)
	go app.Start()
	return app, ramdb
}
