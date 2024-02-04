package repo

import (
	"context"
	"fmt"

	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type TransactionInterface interface {
	SaveTransaction(ctx context.Context, transaction *model.Transaction) error
	GetTransactions(ctx context.Context) ([]model.Transaction, error)
	GetTransactionByHash(ctx context.Context, hash string) (*model.Transaction, error)
}

type Transaction struct {
	db *gorm.DB
	l  hclog.Logger
}

func NewTransaction(db *gorm.DB, l hclog.Logger) *Transaction {
	return &Transaction{db, l}
}

func (repo *Transaction) SaveTransaction(ctx context.Context, transaction *model.Transaction) error {
	if err := repo.db.Create(transaction).Error; err != nil {
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

func (repo *Transaction) GetTransactionByHash(ctx context.Context, hash string) (*model.Transaction, error) {
	var transaction model.Transaction
	var count int64
	repo.db.WithContext(ctx).Model(&transaction).Count(&count)
	if count == 0 {
		return nil, fmt.Errorf("there are no transactions in db")
	}
	if err := repo.db.WithContext(ctx).Where("transaction_hash = ?", hash).First(&transaction).Error; err != nil {
		repo.l.Error("could not find transaction by hash", "hash", hash, "error", err)
		return nil, fmt.Errorf("could not find transaction by hash: %v", err)
	}
	return &transaction, nil
}
