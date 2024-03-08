// internal/repository/account_repository.go
package repository

import (
	"fmt"

	"github.com/kurdilesmana/go-account-api/apps/account/internal/domain"
	"github.com/kurdilesmana/go-account-api/apps/account/pkg/logging"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx *gorm.DB, transaction *domain.Transaction) (ID int, err error)
	CreateDetail(tx *gorm.DB, transactionDetail *domain.TransactionDetail) (err error)
	GetTransactionsByAccountNumber(tx *gorm.DB, AccountNumber string) (Mutations *[]domain.TransactionMutations, err error)
}

type transactionRepository struct {
	db  *gorm.DB
	log *logging.Logger
}

func NewTransactionRepository(db *gorm.DB, log *logging.Logger) TransactionRepository {
	return &transactionRepository{db, log}
}

func (r *transactionRepository) Create(tx *gorm.DB, transaction *domain.Transaction) (ID int, err error) {
	r.log.Info(logrus.Fields{}, transaction, "start create transaction...")
	res := tx.Create(&transaction)
	if res.Error != nil {
		remark := "failed to create transaction"
		r.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, transaction, remark)
		err = fmt.Errorf(remark)
		return
	}

	ID = int(transaction.ID)

	return
}

func (r *transactionRepository) CreateDetail(tx *gorm.DB, transactionDetail *domain.TransactionDetail) (err error) {
	r.log.Info(logrus.Fields{}, transactionDetail, "start create transaction detail...")
	res := tx.Create(&transactionDetail)
	if res.Error != nil {
		remark := "failed to create transaction detail"
		r.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, transactionDetail, remark)
		err = fmt.Errorf(remark)
		return
	}
	return
}

func (r *transactionRepository) GetTransactionsByAccountNumber(tx *gorm.DB, AccountNumber string) (Mutations *[]domain.TransactionMutations, err error) {
	r.log.Info(logrus.Fields{"AccountNumber": AccountNumber}, nil, "get account transaction...")
	res := tx.Table("transaction_details td").
		Select("t.transaction_time, t.transaction_code, td.mutation, td.amount").
		Joins("inner join transactions t on t.id = td.transaction_id").
		Where("td.account_number = ?", AccountNumber).
		Order("t.id desc").
		Find(&Mutations)
	if res.Error != nil {
		remark := "failed to get mutation by account number"
		r.log.Error(logrus.Fields{
			"error": res.Error.Error(),
		}, nil, remark)
		err = fmt.Errorf(remark)
		return
	}

	return
}
