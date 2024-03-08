// internal/initializer/handler_initializer.go
package handler

import (
	"github.com/kurdilesmana/go-account-api/apps/gl/internal/service"
	"github.com/kurdilesmana/go-account-api/apps/gl/pkg/logging"
	"github.com/kurdilesmana/go-account-api/apps/gl/pkg/validator"
)

type Handlers struct {
	JournalHandler JournalHandler
}

func InitHandlers(services *service.Services, log *logging.Logger, validator *validator.RequestValidator) *Handlers {
	return &Handlers{
		JournalHandler: NewJournalHandler(services.JournalService, log, validator),
	}
}
