package repository

import (
	"github.com/kurdilesmana/go-account-api/apps/account/pkg/logging"
	"gorm.io/gorm"
)

type Repositories struct {
	BaseRepository        BaseRepository
	AccountRepository     AccountRepository
	TransactionRepository TransactionRepository
}

func InitRepositories(db *gorm.DB, log *logging.Logger) *Repositories {
	return &Repositories{
		BaseRepository:        NewBaseRepository(db, log),
		AccountRepository:     NewAccountRepository(db, log),
		TransactionRepository: NewTransactionRepository(db, log),
	}
}
