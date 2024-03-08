package service

import (
	"github.com/kurdilesmana/go-account-api/apps/gl/internal/repository"
	"github.com/kurdilesmana/go-account-api/apps/gl/pkg/logging"
)

type Services struct {
	JournalService JournalService
}

func InitServices(repos *repository.Repositories, log *logging.Logger) *Services {
	return &Services{
		JournalService: NewAccountService(repos, log),
	}
}
