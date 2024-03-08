package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		}
		// Lakukan validasi token jika diperlukan
		// contoh: cek token di database atau service lain
		return next(c)
	}
}
