package repository

import (
	"github.com/kurdilesmana/go-account-api/apps/gl/pkg/logging"
	"gorm.io/gorm"
)

type Repositories struct {
	BaseRepository    BaseRepository
	JournalRepository JournalRepository
}

func InitRepositories(db *gorm.DB, log *logging.Logger) *Repositories {
	return &Repositories{
		BaseRepository:    NewBaseRepository(db, log),
		JournalRepository: NewJournalRepository(db, log),
	}
}
