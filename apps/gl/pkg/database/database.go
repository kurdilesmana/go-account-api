// pkg/database/database.go
package database

import (
	"fmt"

	"github.com/kurdilesmana/go-account-api/apps/gl/internal/domain"
	"github.com/kurdilesmana/go-account-api/apps/gl/pkg/logging"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(DB_Host, DB_User, DB_Password, DB_Name string, DB_Port int, log *logging.Logger) *gorm.DB {
	log.Info(logrus.Fields{}, nil, "connecting database...")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", DB_User, DB_Password, DB_Host, DB_Port, DB_Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migrate the Account struct
	log.Info(logrus.Fields{}, nil, "start migrate database...")
	db.AutoMigrate(&domain.Journal{})

	log.Info(logrus.Fields{}, nil, "database connected...")
	return db
}
