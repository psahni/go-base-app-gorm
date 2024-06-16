package db

import (
	"fmt"

	"github.com/go-lang-base-app/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	cfg := config.GetConfig()
	fmt.Println(cfg)
	connectionUrl := cfg.Database.ConnectionUrl
	gormDB, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
