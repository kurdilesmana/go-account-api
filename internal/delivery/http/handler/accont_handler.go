// delivery/account_handler.go
package handler

import (
	"net/http"

	"github.com/kurdilesmana/go-account-api/internal/domain"
	"github.com/kurdilesmana/go-account-api/internal/service"
	"github.com/kurdilesmana/go-account-api/pkg/logging"
	"github.com/kurdilesmana/go-account-api/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AccountHandler interface {
	CreateAccount(c echo.Context) error
}

type accountHandler struct {
	service   service.AccountService
	log       *logging.Logger
	validator *validator.RequestValidator
}

func NewAccountHandler(service service.AccountService, log *logging.Logger, validator *validator.RequestValidator) *accountHandler {
	return &accountHandler{service, log, validator}
}

func (h *accountHandler) CreateAccount(c echo.Context) (err error) {
	h.log.Info(logrus.Fields{}, nil, "Start create account request")
	account := new(domain.CreateAccount)
	if err := c.Bind(account); err != nil {
		remark := "failed to parse request to create account"
		h.log.Error(logrus.Fields{"error": err.Error()}, account, remark)
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, nil, "validate create account request")
	if err := h.validator.Validate(account); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, account, "do create account...")
	AccountNumber, err := h.service.CreateAccount(account)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, account, "done create account.")
	return c.JSON(http.StatusCreated, map[string]string{"account_number": AccountNumber})
}
