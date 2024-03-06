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

type TransactionHandler interface {
	Saving(c echo.Context) error
	CashWithdrawl(c echo.Context) error
	Transfer(c echo.Context) error
}

type transactionHandler struct {
	service   service.TransactionService
	log       *logging.Logger
	validator *validator.RequestValidator
}

func NewTransactionHandler(service service.TransactionService, log *logging.Logger, validator *validator.RequestValidator) *transactionHandler {
	return &transactionHandler{service, log, validator}
}

func (h *transactionHandler) Saving(c echo.Context) (err error) {
	h.log.Info(logrus.Fields{}, nil, "Start saving request")
	saving := new(domain.Saving)
	if err := c.Bind(saving); err != nil {
		remark := "failed to parse request to saving"
		h.log.Error(logrus.Fields{"error": err.Error()}, saving, remark)
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, nil, "validate saving request")
	if err := h.validator.Validate(saving); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, saving, "do saving...")
	Balance, err := h.service.Saving(saving)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, saving, "done saving.")
	return c.JSON(http.StatusCreated, map[string]float64{"balance": Balance})
}

func (h *transactionHandler) CashWithdrawl(c echo.Context) (err error) {
	h.log.Info(logrus.Fields{}, nil, "Start CashWithdrawl request")
	cashwithdrawl := new(domain.CashWithdrawl)
	if err := c.Bind(cashwithdrawl); err != nil {
		remark := "failed to parse request to CashWithdrawl"
		h.log.Error(logrus.Fields{"error": err.Error()}, cashwithdrawl, remark)
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, nil, "validate CashWithdrawl request")
	if err := h.validator.Validate(cashwithdrawl); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, cashwithdrawl, "do CashWithdrawl...")
	Balance, err := h.service.CashWithdrawl(cashwithdrawl)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, cashwithdrawl, "done CashWithdrawl.")
	return c.JSON(http.StatusCreated, map[string]float64{"balance": Balance})
}

func (h *transactionHandler) Transfer(c echo.Context) (err error) {
	h.log.Info(logrus.Fields{}, nil, "Start Transfer request")
	Transfer := new(domain.Transfer)
	if err := c.Bind(Transfer); err != nil {
		remark := "failed to parse request to Transfer"
		h.log.Error(logrus.Fields{"error": err.Error()}, Transfer, remark)
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, nil, "validate Transfer request")
	if err := h.validator.Validate(Transfer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, Transfer, "do Transfer...")
	Balance, err := h.service.Transfer(Transfer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, Transfer, "done Transfer.")
	return c.JSON(http.StatusCreated, map[string]float64{"balance": Balance})
}
