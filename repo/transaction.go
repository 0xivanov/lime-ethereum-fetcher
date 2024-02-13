package repo

import (
	"context"
	"fmt"
	"strconv"

	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/hashicorp/go-hclog"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type TransactionInterface interface {
	SaveTransaction(ctx context.Context, transaction *model.Transaction) error
	GetTransactions(ctx context.Context) ([]model.Transaction, error)
	GetTransactionByHash(ctx context.Context, hash string) (*model.Transaction, error)
	GetTransactionByID(ctx context.Context, id int) (*model.Transaction, error)
	GetTransactionsByUsername(ctx context.Context, username string) ([]model.Transaction, error)
}

type Transaction struct {
	db  *gorm.DB
	rdb *redis.Client
	l   hclog.Logger
}

func NewTransaction(db *gorm.DB, rdb *redis.Client, l hclog.Logger) *Transaction {
	return &Transaction{db, rdb, l}
}

func (repo *Transaction) SaveTransaction(ctx context.Context, transaction *model.Transaction) error {
	if err := repo.db.Create(transaction).Error; err != nil {
		repo.l.Error("could not create transaction", "error", err)
		return fmt.Errorf("could not create transaction: %v", err)
	}

	if err := repo.rdb.SAdd(ctx, transaction.Username, transaction.ID).Err(); err != nil {
		repo.l.Error("could not save the transaction id to redis", "error", err)
		return fmt.Errorf("could not save the transaction id to redis: %v", err)
	}
	repo.l.Info("Stored transaction in Redis", "ID", transaction.ID, "username", transaction.Username)
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

func (repo *Transaction) GetTransactionByID(ctx context.Context, id int) (*model.Transaction, error) {
	var transaction model.Transaction
	var count int64
	repo.db.WithContext(ctx).Model(&transaction).Count(&count)
	if count == 0 {
		return nil, fmt.Errorf("there are no transactions in db")
	}
	if err := repo.db.WithContext(ctx).Where("id = ?", id).First(&transaction).Error; err != nil {
		repo.l.Error("could not find transaction by id", "id", id, "error", err)
		return nil, fmt.Errorf("could not find transaction by id: %v", err)
	}
	return &transaction, nil
}

func (repo *Transaction) GetTransactionsByUsername(ctx context.Context, username string) ([]model.Transaction, error) {
	var transactions []model.Transaction
	transactionIDs, err := repo.rdb.SMembers(ctx, username).Result()
	if err != nil {
		return nil, err
	}

	for _, txnID := range transactionIDs {
		id, err := strconv.Atoi(txnID)
		if err != nil {
			repo.l.Error("failed to convert transaction ID to integer", "transaction_id", txnID, "error", err)
			continue
		}

		transaction, err := repo.GetTransactionByID(ctx, id)
		if err != nil {
			repo.l.Error("failed to retrieve transaction by ID", "transaction_id", id, "error", err)
			continue
		}

		transactions = append(transactions, *transaction)
	}

	return transactions, nil
}
