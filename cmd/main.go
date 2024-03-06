// cmd/main.go
package main

import (
	"github.com/kurdilesmana/go-account-api/internal/delivery/http/handler"
	"github.com/kurdilesmana/go-account-api/internal/delivery/http/router"
	"github.com/kurdilesmana/go-account-api/internal/repository"
	"github.com/kurdilesmana/go-account-api/internal/service"
	"github.com/kurdilesmana/go-account-api/pkg/database"
	"github.com/kurdilesmana/go-account-api/pkg/logging"
	"github.com/kurdilesmana/go-account-api/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()

	// Set Viper to read from .env file
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// APP configuration
	APP_NAME := viper.GetString("APP_NAME")
	APP_PORT := viper.GetString("APP_PORT")

	// DATABASE configuration
	DB_Host := viper.GetString("DB_HOST")
	DB_User := viper.GetString("DB_USER")
	DB_Password := viper.GetString("DB_PASSWORD")
	DB_Name := viper.GetString("DB_DATABASE")
	DB_Port := viper.GetInt("DB_PORT")

	// Dependency injection
	logger := logging.NewLogger(APP_NAME)
	db := database.InitDB(DB_Host, DB_User, DB_Password, DB_Name, DB_Port, logger)
	repos := repository.InitRepositories(db, logger)
	services := service.InitServices(repos, logger)

	// Inisialisasi validator untuk setiap request
	validator := validator.NewRequestValidator()

	// Set middleware untuk validasi request
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("validator", validator)
			return next(c)
		}
	})

	handlers := handler.InitHandlers(services, logger, validator)
	router.InitRoutes(e, handlers)

	e.Logger.Fatal(e.Start(APP_PORT))
}
