// router/router.go
package router

import (
	"github.com/kurdilesmana/go-account-api/internal/delivery/http/handler"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, handlers *handler.Handlers) {
	AccountHandler := handlers.AccountHandler

	e.POST("/daftar", AccountHandler.CreateAccount)
}
