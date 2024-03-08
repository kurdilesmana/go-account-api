// internal/repository/account_repository.go
package repository

import (
	"fmt"

	"github.com/kurdilesmana/go-account-api/apps/gl/internal/domain"
	"github.com/kurdilesmana/go-account-api/apps/gl/pkg/logging"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JournalRepository interface {
	Create(tx *gorm.DB, account *domain.Journal) error
}

type journalRepository struct {
	db  *gorm.DB
	log *logging.Logger
}

func NewJournalRepository(db *gorm.DB, log *logging.Logger) JournalRepository {
	return &journalRepository{db, log}
}

func (r *journalRepository) Create(tx *gorm.DB, journal *domain.Journal) (err error) {
	r.log.Info(logrus.Fields{}, journal, "start create journal...")
	res := tx.Create(&journal)
	if res.Error != nil {
		remark := "failed to create rekening"
		r.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, journal, remark)
		err = fmt.Errorf(remark)
		return
	}
	return
}
