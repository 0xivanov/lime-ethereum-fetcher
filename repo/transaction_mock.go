package repo

import (
	"context"
	"fmt"

	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
)

type MockTransaction struct {
}

func NewMockTransaction() *MockTransaction {
	return &MockTransaction{}
}

func (m *MockTransaction) SaveTransactions(ctx context.Context, transactions *[]model.Transaction) error {
	return nil
}

func (m *MockTransaction) GetTransactions(ctx context.Context) ([]model.Transaction, error) {
	// Mocked list of transactions
	transactions := []model.Transaction{
		{
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
		},
		{
			ID:                2,
			TransactionHash:   "mocked hash 2",
			TransactionStatus: "mocked status 2",
			BlockHash:         "mocked block hash 2",
			BlockNumber:       "mocked block number 2",
			From:              "mocked from address 2",
			To:                "mocked to address 2",
			ContractAddress:   "mocked contract address 2",
			LogsCount:         "mocked logs count 2",
			Input:             "mocked input 2",
			Value:             "mocked value 2",
		},
	}

	return transactions, nil

}

type MockTransactionError struct {
}

func NewMockTransactionError() *MockTransactionError {
	return &MockTransactionError{}
}

func (m *MockTransactionError) SaveTransactions(ctx context.Context, transactions *[]model.Transaction) error {
	return nil
}

func (m *MockTransactionError) GetTransactions(ctx context.Context) ([]model.Transaction, error) {
	return nil, fmt.Errorf("error")
}
