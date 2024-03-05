// internal/repository/account_repository.go
package repository

import (
	"fmt"

	"github.com/kurdilesmana/go-account-api/internal/domain"
	"github.com/kurdilesmana/go-account-api/pkg/logging"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(tx *gorm.DB, account *domain.Account) error
	CheckAccountExist(tx *gorm.DB, NIK, PhoneNumber string) (exist bool, err error)
}

type accountRepository struct {
	db  *gorm.DB
	log *logging.Logger
}

func NewAccountRepository(db *gorm.DB, log *logging.Logger) AccountRepository {
	return &accountRepository{db, log}
}

func (t *accountRepository) Begin() (tx *gorm.DB) {
	return t.db.Begin()
}

func (t *accountRepository) Rollback(tx *gorm.DB) {
	tx.Rollback()
}

func (t *accountRepository) Commit(tx *gorm.DB) {
	tx.Commit()
}

func (r *accountRepository) Create(tx *gorm.DB, account *domain.Account) (err error) {
	r.log.Info(logrus.Fields{}, account, "start create account...")
	res := tx.Create(&account)
	if res.Error != nil {
		remark := "failed to create rekening"
		r.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}
	return
}

func (r *accountRepository) CheckAccountExist(tx *gorm.DB, NIK, PhoneNumber string) (exist bool, err error) {
	r.log.Info(logrus.Fields{"NIK": NIK, "PhoneNumber": PhoneNumber}, nil, "check account exist...")
	var account []domain.Account
	res := tx.Where(domain.Account{NIK: NIK}).Where(domain.Account{PhoneNumber: PhoneNumber}).Find(&account)
	if res.Error != nil {
		remark := "failed to get account by nik phonenumber"
		r.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}

	if res.RowsAffected > 0 {
		exist = true
	}
	return
}
