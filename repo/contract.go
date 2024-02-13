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
	GetPersons(ctx context.Context) ([]model.PersonInfoEvent, error)
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
		repo.l.Error("could not create person event", "error", err)
		return fmt.Errorf("could not create person event: %v", err)
	}

	return nil
}

func (repo *Contract) GetPersons(ctx context.Context) ([]model.PersonInfoEvent, error) {
	var persons []model.PersonInfoEvent
	if err := repo.db.Find(&persons).Error; err != nil {
		repo.l.Error("could not retrieve persons", "error", err)
		return nil, fmt.Errorf("could not retrieve persons: %v", err)
	}
	return persons, nil
}
