// internal/service/account_service.go
package service

import (
	"fmt"
	"time"

	"github.com/kurdilesmana/go-account-api/internal/domain"
	"github.com/kurdilesmana/go-account-api/internal/repository"
	"github.com/kurdilesmana/go-account-api/pkg/logging"
	"github.com/sirupsen/logrus"
)

type TransactionService interface {
	Saving(account *domain.Saving) (Balance float64, err error)
	CashWithdrawl(account *domain.CashWithdrawl) (Balance float64, err error)
	Transfer(balanceInquiry *domain.Transfer) (Balance float64, err error)
}

type transactionService struct {
	repo *repository.Repositories
	log  *logging.Logger
}

func NewTransactionService(repo *repository.Repositories, log *logging.Logger) TransactionService {
	return &transactionService{repo, log}
}

func (s *transactionService) Saving(params *domain.Saving) (Balance float64, err error) {
	// initiate transaction
	tx := s.repo.BaseRepository.Begin()

	// validate account exist
	isExist, err := s.repo.AccountRepository.CheckAccountExistByAccountNUmber(tx, params.AccountNumber)
	if err != nil {
		err = fmt.Errorf("failed to check exist account")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	if !isExist {
		err = fmt.Errorf("failed to saving, account not exist")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// create transaction
	transaction := domain.Transaction{
		TransactionTime: time.Now(),
		TransactionCode: "C",
		Amount:          params.Amount,
	}
	IDTransaction, err := s.repo.TransactionRepository.Create(tx, &transaction)
	if err != nil {
		err = fmt.Errorf("failed to create transaction")
		s.log.Warn(logrus.Fields{}, transaction, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// create transaction detail
	transactionDetail := domain.TransactionDetail{
		TransactionID: IDTransaction,
		Mutation:      "C",
		AccountNumber: params.AccountNumber,
		Amount:        params.Amount,
	}
	if err = s.repo.TransactionRepository.CreateDetail(tx, &transactionDetail); err != nil {
		err = fmt.Errorf("failed to create transaction")
		s.log.Warn(logrus.Fields{}, transaction, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// update balance account
	Balance, err = s.repo.AccountRepository.UpdateBalanceAccount(tx, params.AccountNumber, params.Amount)
	if err != nil {
		err = fmt.Errorf("failed to update balance account")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// commit transacton
	s.repo.BaseRepository.Commit(tx)

	return
}

func (s *transactionService) CashWithdrawl(params *domain.CashWithdrawl) (Balance float64, err error) {
	// initiate transaction
	tx := s.repo.BaseRepository.Begin()

	// validate account exist
	isExist, err := s.repo.AccountRepository.CheckAccountExistByAccountNUmber(tx, params.AccountNumber)
	if err != nil {
		err = fmt.Errorf("failed to check exist account")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	if !isExist {
		err = fmt.Errorf("failed to cashwithdrawl, account not exist")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// Get Balance by account number
	BalanceExisting, err := s.repo.AccountRepository.GetBalanceByAccountNumber(tx, params.AccountNumber)
	if err != nil {
		err = fmt.Errorf("failed to get balance account")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// validate balance
	if (BalanceExisting - params.Amount) < 0.0 {
		err = fmt.Errorf("failed to cashwithdrawl, insufficient balance")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// create transaction
	transaction := domain.Transaction{
		TransactionTime: time.Now(),
		TransactionCode: "D",
		Amount:          params.Amount,
	}
	IDTransaction, err := s.repo.TransactionRepository.Create(tx, &transaction)
	if err != nil {
		err = fmt.Errorf("failed to create transaction")
		s.log.Warn(logrus.Fields{}, transaction, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// create transaction detail
	transactionDetail := domain.TransactionDetail{
		TransactionID: IDTransaction,
		Mutation:      "D",
		AccountNumber: params.AccountNumber,
		Amount:        params.Amount,
	}
	if err = s.repo.TransactionRepository.CreateDetail(tx, &transactionDetail); err != nil {
		err = fmt.Errorf("failed to create transaction")
		s.log.Warn(logrus.Fields{}, transaction, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// update balance account
	params.Amount = params.Amount * -1
	Balance, err = s.repo.AccountRepository.UpdateBalanceAccount(tx, params.AccountNumber, params.Amount)
	if err != nil {
		err = fmt.Errorf("failed to update balance account")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// commit transacton
	s.repo.BaseRepository.Commit(tx)

	return
}

func (s *transactionService) Transfer(params *domain.Transfer) (Balance float64, err error) {
	// initiate transaction
	tx := s.repo.BaseRepository.Begin()

	// validate source account exist
	isExist, err := s.repo.AccountRepository.CheckAccountExistByAccountNUmber(tx, params.SrcAccountNumber)
	if err != nil {
		err = fmt.Errorf("failed to check exist account")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	if !isExist {
		err = fmt.Errorf("failed to transfer, source account not exist")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// validate destination account exist
	isExist, err = s.repo.AccountRepository.CheckAccountExistByAccountNUmber(tx, params.DestAccountNumber)
	if err != nil {
		err = fmt.Errorf("failed to check exist account")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	if !isExist {
		err = fmt.Errorf("failed to transfer, destination account not exist")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// Get Balance by account number
	BalanceExisting, err := s.repo.AccountRepository.GetBalanceByAccountNumber(tx, params.SrcAccountNumber)
	if err != nil {
		err = fmt.Errorf("failed to get balance source account")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// validate balance
	if (BalanceExisting - params.Amount) < 0.0 {
		err = fmt.Errorf("failed to cashwithdrawl, insufficient balance")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// create transaction
	transaction := domain.Transaction{
		TransactionTime: time.Now(),
		TransactionCode: "T",
		Amount:          params.Amount,
	}
	IDTransaction, err := s.repo.TransactionRepository.Create(tx, &transaction)
	if err != nil {
		err = fmt.Errorf("failed to create transaction")
		s.log.Warn(logrus.Fields{}, transaction, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// create transaction detail src
	transactionDetailSrc := domain.TransactionDetail{
		TransactionID: IDTransaction,
		Mutation:      "D",
		AccountNumber: params.SrcAccountNumber,
		Amount:        params.Amount,
	}
	if err = s.repo.TransactionRepository.CreateDetail(tx, &transactionDetailSrc); err != nil {
		err = fmt.Errorf("failed to create transaction")
		s.log.Warn(logrus.Fields{}, transaction, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// create transaction detail dest
	transactionDetailDest := domain.TransactionDetail{
		TransactionID: IDTransaction,
		Mutation:      "C",
		AccountNumber: params.DestAccountNumber,
		Amount:        params.Amount,
	}
	if err = s.repo.TransactionRepository.CreateDetail(tx, &transactionDetailDest); err != nil {
		err = fmt.Errorf("failed to create transaction")
		s.log.Warn(logrus.Fields{}, transaction, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// update balance account src
	Balance, err = s.repo.AccountRepository.UpdateBalanceAccount(tx, params.SrcAccountNumber, params.Amount*-1)
	if err != nil {
		err = fmt.Errorf("failed to update balance account")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// update balance account dst
	_, err = s.repo.AccountRepository.UpdateBalanceAccount(tx, params.DestAccountNumber, params.Amount*1)
	if err != nil {
		err = fmt.Errorf("failed to update balance account")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// commit transacton
	s.repo.BaseRepository.Commit(tx)

	return
}
