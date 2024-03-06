package service

import (
	"github.com/kurdilesmana/go-account-api/internal/repository"
	"github.com/kurdilesmana/go-account-api/pkg/logging"
)

type Services struct {
	AccountService     AccountService
	TransactionService TransactionService
}

func InitServices(repos *repository.Repositories, log *logging.Logger) *Services {
	return &Services{
		AccountService:     NewAccountService(repos, log),
		TransactionService: NewTransactionService(repos, log),
	}
}
