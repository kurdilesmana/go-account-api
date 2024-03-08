// router/router.go
package router

import (
	"github.com/kurdilesmana/go-account-api/apps/gl/internal/delivery/http/handler"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, handlers *handler.Handlers) {
	JournalHandler := handlers.JournalHandler

	e.POST("/daftar", JournalHandler.CreateJournal)
}
