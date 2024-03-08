// internal/repository/account_repository.go
package repository

import (
	"github.com/kurdilesmana/go-account-api/apps/account/pkg/logging"
	"gorm.io/gorm"
)

type BaseRepository interface {
	Begin() (tx *gorm.DB)
	Rollback(tx *gorm.DB)
	Commit(tx *gorm.DB)
}

type baseRepository struct {
	db  *gorm.DB
	log *logging.Logger
}

func NewBaseRepository(db *gorm.DB, log *logging.Logger) BaseRepository {
	return &baseRepository{db, log}
}

func (b *baseRepository) Begin() (tx *gorm.DB) {
	return b.db.Begin()
}

func (b *baseRepository) Rollback(tx *gorm.DB) {
	tx.Rollback()
}

func (b *baseRepository) Commit(tx *gorm.DB) {
	tx.Commit()
}
