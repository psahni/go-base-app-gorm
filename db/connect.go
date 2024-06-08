package db

import (
	"github.com/go-lang-base-app/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	connectionUrl := config.GetConfig().Database.ConnectionUrl
	gormDB, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
