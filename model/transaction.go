package model

type Transaction struct {
	ID                int    `json:"-" gorm:"primaryKey"`
	TransactionHash   string `json:"hash" gorm:"column:transaction_hash"`
	TransactionStatus string `json:"type" gorm:"column:transaction_status"`
	BlockHash         string `json:"blockHash" gorm:"column:block_hash"`
	BlockNumber       string `json:"blockNumber" gorm:"column:block_number"`
	From              string `json:"from" gorm:"column:from_address"`
	To                string `json:"to" gorm:"column:to_address"`
	ContractAddress   string `json:"-" gorm:"column:contract_address"`
	LogsCount         string `json:"-" gorm:"column:logs_count"`
	Input             string `json:"input" gorm:"column:input"`
	Value             string `json:"value" gorm:"column:value"`
	Username          string `json:"-" gorm:"-"`
}

type TransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
}
