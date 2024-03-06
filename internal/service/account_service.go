// internal/service/account_service.go
package service

import (
	"fmt"

	"github.com/kurdilesmana/go-account-api/internal/domain"
	"github.com/kurdilesmana/go-account-api/internal/repository"
	"github.com/kurdilesmana/go-account-api/pkg/logging"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AccountService interface {
	CreateAccount(account *domain.CreateAccount) (AccountNumber string, err error)
	BalanceInquiry(balanceInquiry *domain.BalanceInquiry) (Balance float64, err error)
	TransactionInquiry(transactionInquiry *domain.TransactionInquiry) (Mutations *[]domain.TransactionMutations, err error)
}

type accountService struct {
	repo *repository.Repositories
	log  *logging.Logger
}

func NewAccountService(repo *repository.Repositories, log *logging.Logger) AccountService {
	return &accountService{repo, log}
}

func (s *accountService) CreateAccount(createAccount *domain.CreateAccount) (AccountNumber string, err error) {
	// initiate transaction
	tx := s.repo.BaseRepository.Begin()

	// validate account exist
	isExist, err := s.repo.AccountRepository.CheckAccountExist(tx, createAccount.NIK, createAccount.PhoneNumber)
	if err != nil {
		err = fmt.Errorf("failed to check exist rekening")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	if isExist {
		err = fmt.Errorf("failed to create, rekening exist")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// Generate an account number
	AccountNumber = createAccount.PhoneNumber + "0001"

	// Hash PIN
	hashedPIN, err := bcrypt.GenerateFromPassword([]byte(createAccount.PIN), bcrypt.MinCost)
	if err != nil {
		err = fmt.Errorf("failed to hashed PIN")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// create account
	account := domain.Account{
		AccountNumber: AccountNumber,
		Name:          createAccount.Name,
		NIK:           createAccount.NIK,
		PhoneNumber:   createAccount.PhoneNumber,
		PIN:           string(hashedPIN),
		Balance:       0.0,
	}
	if err = s.repo.AccountRepository.Create(tx, &account); err != nil {
		err = fmt.Errorf("failed to create rekening")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// commit transacton
	s.repo.BaseRepository.Commit(tx)

	return
}

func (s *accountService) BalanceInquiry(balanceInquiry *domain.BalanceInquiry) (Balance float64, err error) {
	// initiate transaction
	tx := s.repo.BaseRepository.Begin()

	// validate account exist
	isExist, err := s.repo.AccountRepository.CheckAccountExistByAccountNUmber(tx, balanceInquiry.AccountNumber)
	if err != nil {
		err = fmt.Errorf("failed to check exist rekening")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	if !isExist {
		err = fmt.Errorf("failed to create, rekening not exist")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// Get Balance by account number
	Balance, err = s.repo.AccountRepository.GetBalanceByAccountNumber(tx, balanceInquiry.AccountNumber)
	if err != nil {
		err = fmt.Errorf("failed to get balance account")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// commit transacton
	s.repo.BaseRepository.Commit(tx)

	return
}

func (s *accountService) TransactionInquiry(transactionInquiry *domain.TransactionInquiry) (Mutations *[]domain.TransactionMutations, err error) {
	// initiate transaction
	tx := s.repo.BaseRepository.Begin()

	// validate account exist
	isExist, err := s.repo.AccountRepository.CheckAccountExistByAccountNUmber(tx, transactionInquiry.AccountNumber)
	if err != nil {
		err = fmt.Errorf("failed to check exist rekening")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	if !isExist {
		err = fmt.Errorf("failed to create, rekening not exist")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// Get Balance by account number
	Mutations, err = s.repo.TransactionRepository.GetTransactionsByAccountNumber(tx, transactionInquiry.AccountNumber)
	if err != nil {
		err = fmt.Errorf("failed to get balance account")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// commit transacton
	s.repo.BaseRepository.Commit(tx)

	return
}
