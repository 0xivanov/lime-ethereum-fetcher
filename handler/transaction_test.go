package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	mock_repo "github.com/0xivanov/lime-ethereum-fetcher-go/mocks"
	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

/*
*

	Unit tests

*
*/
func TestFetchTransactionsFromEthereum_Positive(t *testing.T) {
	LoadEnvVars(t)
	subject := &Transaction{hclog.Default(), nil, os.Getenv("ETH_NODE_URL")}

	var transactionHashes = "0x518e92e52a79128998ebda10e7db80798045a4268237d50488e4dbdcc5da2986"
	result, err := subject.fetchTransactionFromEthereum(transactionHashes)
	if err != nil {
		t.Errorf("fetchTransactionsFromEthereum returned error: %v", err)
	}
	fmt.Println(result)
}

func TestFetchTransactionsFromEthereum_Negative(t *testing.T) {
	LoadEnvVars(t)
	subject := &Transaction{hclog.Default(), nil, os.Getenv("ETH_NODE_URL")}

	var transactionHash = "asdf"
	_, err := subject.fetchTransactionFromEthereum(transactionHash)

	expectedErrorMessage := "unexpected status code while fetching transactions from ethereum: 400"
	if err != nil && !strings.Contains(err.Error(), expectedErrorMessage) {
		t.Errorf("error message does not contain expected substring '%s'", expectedErrorMessage)
	}
}

func TestGetTransactions_Positive(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	router := gin.Default()
	mockRepo := mock_repo.NewMockTransactionInterface(ctrl)
	mockRepo.EXPECT().GetTransactions(gomock.Any()).Return([]model.Transaction{{
		ID:                1,
		TransactionHash:   "mocked hash 1",
		TransactionStatus: "mocked status 1",
		BlockHash:         "mocked block hash 1",
		BlockNumber:       "mocked block number 1",
		From:              "mocked from address 1",
		To:                "mocked to address 1",
		ContractAddress:   "mocked contract address 1",
		LogsCount:         "mocked logs count 1",
		Input:             "mocked input 1",
		Value:             "mocked value 1",
	}}, nil)
	subject := &Transaction{hclog.Default(), mockRepo, ""}

	// when
	router.GET("/lime/all", subject.GetTransactions)
	req, err := http.NewRequest("GET", "/lime/all", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// then
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetTransactions_NegativeDbError(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	router := gin.Default()
	mockRepo := mock_repo.NewMockTransactionInterface(ctrl)
	mockRepo.EXPECT().GetTransactions(gomock.Any()).Return(nil, errors.New("mocked error"))
	subject := &Transaction{hclog.Default(), mockRepo, ""}

	// when
	router.GET("/lime/all", subject.GetTransactions)
	req, err := http.NewRequest("GET", "/lime/all", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// then
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func LoadEnvVars(t *testing.T) {
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
}
