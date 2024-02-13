package repo

import (
	"context"
	"fmt"

	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
)

type ContractInterface interface {
	SavePersonInfoUpdatedEvent(ctx context.Context, personEvent *model.PersonInfoEvent) error
}

type Contract struct {
	db *gorm.DB
	l  hclog.Logger
}

func NewContract(db *gorm.DB, l hclog.Logger) *Contract {
	return &Contract{db, l}
}

func (repo *Contract) SavePersonInfoUpdatedEvent(ctx context.Context, personEvent *model.PersonInfoEvent) error {
	if err := repo.db.Create(personEvent).Error; err != nil {
		repo.l.Error("could not create person event: %v", err)
		return fmt.Errorf("could not create person event: %v", err)
	}

	return nil
}
