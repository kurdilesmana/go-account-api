// delivery/account_handler.go
package handler

import (
	"net/http"

	"github.com/kurdilesmana/go-account-api/apps/account/internal/domain"
	"github.com/kurdilesmana/go-account-api/apps/account/internal/service"
	"github.com/kurdilesmana/go-account-api/apps/account/pkg/logging"
	"github.com/kurdilesmana/go-account-api/apps/account/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AccountHandler interface {
	CreateAccount(c echo.Context) error
	BalanceInquiry(c echo.Context) error
	Mutation(c echo.Context) error
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

func (h *accountHandler) BalanceInquiry(c echo.Context) (err error) {
	h.log.Info(logrus.Fields{}, nil, "Start balance inquiry request")
	balanceInquiry := &domain.BalanceInquiry{AccountNumber: c.Param("account_number")}

	h.log.Info(logrus.Fields{}, balanceInquiry, "validate balance inquiry request")
	if err := h.validator.Validate(balanceInquiry); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, balanceInquiry, "do balance inquiry...")
	Balance, err := h.service.BalanceInquiry(balanceInquiry)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, balanceInquiry, "done balance inquiry.")
	return c.JSON(http.StatusCreated, map[string]float64{"balance": Balance})
}

func (h *accountHandler) Mutation(c echo.Context) (err error) {
	h.log.Info(logrus.Fields{}, nil, "Start mutation request")
	transactionInquiry := &domain.TransactionInquiry{AccountNumber: c.Param("account_number")}

	h.log.Info(logrus.Fields{}, transactionInquiry, "validate mutation request")
	if err := h.validator.Validate(transactionInquiry); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, transactionInquiry, "do mutation...")
	mutation, err := h.service.TransactionInquiry(transactionInquiry)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	h.log.Info(logrus.Fields{}, transactionInquiry, "done mutation.")
	return c.JSON(http.StatusCreated, map[string]interface{}{"mutation": mutation})
}
