package db

import (
	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabse(dialector gorm.Dialector) (*Database, error) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Transaction{})

	return &Database{db: db}, nil
}

func (d *Database) GetDb() *gorm.DB {
	return d.db
}
