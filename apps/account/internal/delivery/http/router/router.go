// router/router.go
package router

import (
	"github.com/kurdilesmana/go-account-api/apps/account/internal/delivery/http/handler"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, handlers *handler.Handlers) {
	AccountHandler := handlers.AccountHandler
	TransactionHandler := handlers.TransactionHandler

	e.POST("/daftar", AccountHandler.CreateAccount)
	e.POST("/tabung", TransactionHandler.Saving)
	e.POST("/tarik", TransactionHandler.CashWithdrawl)
	e.POST("/transfer", TransactionHandler.Transfer)
	e.GET("/saldo/:account_number", AccountHandler.BalanceInquiry)
	e.GET("/mutasi/:account_number", AccountHandler.Mutation)
}
