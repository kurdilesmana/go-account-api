// delivery/account_handler.go
package handler

import (
	"net/http"

	"github.com/kurdilesmana/go-account-api/apps/gl/internal/domain"
	"github.com/kurdilesmana/go-account-api/apps/gl/internal/service"
	"github.com/kurdilesmana/go-account-api/apps/gl/pkg/logging"
	"github.com/kurdilesmana/go-account-api/apps/gl/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type JournalHandler interface {
	CreateJournal(c echo.Context) error
}

type journalHandler struct {
	service   service.JournalService
	log       *logging.Logger
	validator *validator.RequestValidator
}

func NewJournalHandler(service service.JournalService, log *logging.Logger, validator *validator.RequestValidator) *journalHandler {
	return &journalHandler{service, log, validator}
}

func (h *journalHandler) CreateJournal(c echo.Context) (err error) {
	h.log.Info(logrus.Fields{}, nil, "Start create journal request")
	journal := new(domain.CreateJournal)
	if err := c.Bind(journal); err != nil {
		remark := "failed to parse request to create journal"
		h.log.Error(logrus.Fields{"error": err.Error()}, journal, remark)
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, nil, "validate create journal request")
	if err := h.validator.Validate(journal); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, journal, "do create journal...")
	err = h.service.CreateJournal(journal)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, journal, "done create journal.")
	return c.JSON(http.StatusCreated, map[string]string{"message": "OK"})
}
