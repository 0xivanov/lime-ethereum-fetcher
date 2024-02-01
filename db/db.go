package db

import (
	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabse(dbConnectionString string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Transaction{})

	return &Database{db: db}, nil
}

func (d *Database) GetDb() *gorm.DB {
	return d.db
}
