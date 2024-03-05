package repository

import (
	"github.com/kurdilesmana/go-account-api/pkg/logging"
	"gorm.io/gorm"
)

type Repositories struct {
	BaseRepository    BaseRepository
	AccountRepository AccountRepository
}

func InitRepositories(db *gorm.DB, log *logging.Logger) *Repositories {
	return &Repositories{
		BaseRepository:    NewBaseRepository(db, log),
		AccountRepository: NewAccountRepository(db, log),
	}
}