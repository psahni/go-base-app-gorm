package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DB_URL = "postgres://postgres:root@localhost:5432/test_database"

func Connect() (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
