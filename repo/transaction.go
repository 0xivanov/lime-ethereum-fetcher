package repo

import (
	"context"
	"fmt"

	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type TransactionInterface interface {
	SaveTransactions(ctx context.Context, transactions *[]model.Transaction) error
	GetTransactions(ctx context.Context) ([]model.Transaction, error)
}

type Transaction struct {
	db *gorm.DB
	l  hclog.Logger
}

func NewTransaction(db *gorm.DB, l hclog.Logger) *Transaction {
	return &Transaction{db, l}
}

func (repo *Transaction) SaveTransactions(ctx context.Context, transactions *[]model.Transaction) error {
	if err := repo.db.Create(transactions).Error; err != nil {
		repo.l.Error("could not create transaction: %v", err)
		return fmt.Errorf("could not create transaction: %v", err)
	}
	return nil
}

func (repo *Transaction) GetTransactions(ctx context.Context) ([]model.Transaction, error) {
	var transactions []model.Transaction
	if err := repo.db.WithContext(ctx).Find(&transactions).Error; err != nil {
		repo.l.Error("could not find transactions: %v", err)
		return nil, fmt.Errorf("could not find transactions: %v", err)
	}
	return transactions, nil
}
