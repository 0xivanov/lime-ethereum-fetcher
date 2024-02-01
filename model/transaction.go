package model

type Transaction struct {
	ID                int    `json:"id,omitempty" gorm:"primaryKey"`
	TransactionHash   string `json:"transactionHash" gorm:"column:transaction_hash"`
	TransactionStatus int    `json:"transactionStatus" gorm:"column:transaction_status"`
	BlockHash         string `json:"blockHash" gorm:"column:block_hash"`
	BlockNumber       int    `json:"blockNumber" gorm:"column:block_number"`
	From              string `json:"from" gorm:"column:from_address"`
	To                string `json:"to,omitempty" gorm:"column:to_address"`
	ContractAddress   string `json:"contractAddress,omitempty" gorm:"column:contract_address"`
	LogsCount         int    `json:"logsCount" gorm:"column:logs_count"`
	Input             string `json:"input" gorm:"column:input"`
	Value             int    `json:"value" gorm:"column:value"`
}

type TransactionResponse struct {
	Transactions []Transaction `json:"transactions"`
}
