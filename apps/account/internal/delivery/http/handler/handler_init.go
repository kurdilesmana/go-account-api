// internal/initializer/handler_initializer.go
package handler

import (
	"github.com/kurdilesmana/go-account-api/apps/account/internal/service"
	"github.com/kurdilesmana/go-account-api/apps/account/pkg/logging"
	"github.com/kurdilesmana/go-account-api/apps/account/pkg/validator"
)

type Handlers struct {
	AccountHandler     AccountHandler
	TransactionHandler TransactionHandler
}

func InitHandlers(services *service.Services, log *logging.Logger, validator *validator.RequestValidator) *Handlers {
	return &Handlers{
		AccountHandler:     NewAccountHandler(services.AccountService, log, validator),
		TransactionHandler: NewTransactionHandler(services.TransactionService, log, validator),
	}
}
