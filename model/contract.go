package model

type PersonInfo struct {
	Name string
	Age  int64
}

type PersonInfoEvent struct {
	TxHash string `gorm:"not null"`
	Index  uint64 `gorm:"not null"`
	Name   string `gorm:"not null"`
	Age    uint64 `gorm:"not null"`
}
