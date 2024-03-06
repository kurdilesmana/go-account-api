// internal/repository/account_repository.go
package repository

import (
	"fmt"

	"github.com/kurdilesmana/go-account-api/internal/domain"
	"github.com/kurdilesmana/go-account-api/pkg/logging"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountRepository interface {
	Create(tx *gorm.DB, account *domain.Account) error
	CheckAccountExist(tx *gorm.DB, NIK, PhoneNumber string) (exist bool, err error)
	CheckAccountExistByAccountNUmber(tx *gorm.DB, AccountNumber string) (exist bool, err error)
	UpdateBalanceAccount(tx *gorm.DB, AccountNumber string, Amount float64) (Balance float64, err error)
	GetBalanceByAccountNumber(tx *gorm.DB, AccountNumber string) (Balance float64, err error)
}

type accountRepository struct {
	db  *gorm.DB
	log *logging.Logger
}

func NewAccountRepository(db *gorm.DB, log *logging.Logger) AccountRepository {
	return &accountRepository{db, log}
}

func (r *accountRepository) Create(tx *gorm.DB, account *domain.Account) (err error) {
	r.log.Info(logrus.Fields{}, account, "start create account...")
	res := tx.Create(&account)
	if res.Error != nil {
		remark := "failed to create rekening"
		r.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, account, remark)
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

func (r *accountRepository) CheckAccountExistByAccountNUmber(tx *gorm.DB, AccountNumber string) (exist bool, err error) {
	r.log.Info(logrus.Fields{"AccountNumber": AccountNumber}, nil, "check account exist...")
	var account []domain.Account
	res := tx.Where(domain.Account{AccountNumber: AccountNumber}).Find(&account)
	if res.Error != nil {
		remark := "failed to get account by account number"
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

func (t *accountRepository) UpdateBalanceAccount(tx *gorm.DB, AccountNumber string, Amount float64) (Balance float64, err error) {
	var account domain.Account
	res := tx.Model(&account).Clauses(clause.Returning{Columns: []clause.Column{{Name: "balance"}}}).Where(
		"account_number = ?", AccountNumber).UpdateColumn("balance", gorm.Expr("balance + ?", Amount))
	if res.Error != nil {
		remark := "failed to update balance account"
		t.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}
	Balance = account.Balance

	return
}

func (r *accountRepository) GetBalanceByAccountNumber(tx *gorm.DB, AccountNumber string) (Balance float64, err error) {
	r.log.Info(logrus.Fields{"AccountNumber": AccountNumber}, nil, "check account exist...")
	var account domain.Account
	res := tx.Model(&account).Select("balance").Where("account_number = ?", AccountNumber).Scan(&Balance)
	if res.Error != nil {
		remark := "failed to get account by account number"
		r.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}

	return
}
