// internal/service/account_service.go
package service

import (
	"fmt"

	"github.com/kurdilesmana/go-account-api/apps/gl/internal/domain"
	"github.com/kurdilesmana/go-account-api/apps/gl/internal/repository"
	"github.com/kurdilesmana/go-account-api/apps/gl/pkg/logging"
	"github.com/sirupsen/logrus"
)

type JournalService interface {
	CreateJournal(account *domain.CreateJournal) (err error)
}

type journalService struct {
	repo *repository.Repositories
	log  *logging.Logger
}

func NewAccountService(repo *repository.Repositories, log *logging.Logger) JournalService {
	return &journalService{repo, log}
}

func (s *journalService) CreateJournal(createjournal *domain.CreateJournal) (err error) {
	// initiate transaction
	tx := s.repo.BaseRepository.Begin()

	// create account
	journal := domain.Journal{
		TransactionDate: createjournal.TransactionDate,
		CreditAccount:   createjournal.CreditAccount,
		DebetAccount:    createjournal.DebetAccount,
		CreditAmount:    createjournal.CreditAmount,
		DebetAmount:     createjournal.DebetAmount,
	}
	if err = s.repo.JournalRepository.Create(tx, &journal); err != nil {
		err = fmt.Errorf("failed to create rekening")
		s.log.Warn(logrus.Fields{}, nil, err.Error())
		s.repo.BaseRepository.Rollback(tx)
		return
	}

	// commit transacton
	s.repo.BaseRepository.Commit(tx)

	return
}
