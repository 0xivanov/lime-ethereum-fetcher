package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/0xivanov/lime-ethereum-fetcher-go/repo"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

type Transaction struct {
	l          hclog.Logger
	tr         repo.TransactionInterface
	ethNodeUrl string
}

func NewTransaction(l hclog.Logger, tr repo.TransactionInterface) *Transaction {
	ethNodeUrl := os.Getenv("ETH_NODE_URL")
	return &Transaction{l, tr, ethNodeUrl}
}

func (t *Transaction) GetTransactionsWithHashes(ctx *gin.Context) {
	transactionHashes := ctx.QueryArray("transactionHashes")
	t.l.Debug("transaction hashes:", "transactionHashes", transactionHashes)
	var transactions = []model.Transaction{}

	for _, hash := range transactionHashes {
		// try to get transaction from db
		transactionFromDb, err := t.tr.GetTransactionByHash(ctx, hash)
		if transactionFromDb != nil {
			transactions = append(transactions, *transactionFromDb)
			continue
		}
		if err != nil {
			t.l.Info("could not get transaction from db with hash", "hash", hash)
			t.l.Info("trying to get transaction from ethereum", "hash", hash)
		}
		// get transactions from ethereum
		transactionFromEth, err := t.fetchTransactionFromEthereum(hash)
		if err != nil {
			t.l.Error("could not get transactions ethereum", "error", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		t.l.Debug("fetched transactions from ethereum", "transactions", transactions)
		transactions = append(transactions, *transactionFromEth)
		// save to db
		err = t.tr.SaveTransaction(ctx, transactionFromEth)
		if err != nil {
			t.l.Error("could not save transactions", "error", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	t.l.Info("Retrieved transactions", "number of transactions", len(transactions))
	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func (t *Transaction) GetTransactions(ctx *gin.Context) {
	// get transactions from db
	transactions, err := t.tr.GetTransactions(ctx)
	if err != nil {
		t.l.Error("could not get transactions from db", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	t.l.Info("Retrieved transactions", "number of transactions", len(transactions))
	ctx.JSON(http.StatusOK, transactions)
}

func (t *Transaction) fetchTransactionFromEthereum(hash string) (*model.Transaction, error) {
	client := &http.Client{}

	var response struct {
		Result model.Transaction `json:"result"`
	}
	// create a JSON payload
	payload := map[string]interface{}{
		"id":      1,
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionByHash",
		"params":  []interface{}{hash},
	}

	// marshal the payload into JSON
	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", t.ethNodeUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code while fetching transactions from ethereum: %d. request: %v", resp.StatusCode, req)
	}
	// Decode the response JSON
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	return &response.Result, nil
}
