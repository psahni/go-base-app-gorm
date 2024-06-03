package migration

import "gorm.io/gorm"

type Direction int

const (
	UP Direction = iota
	DOWN
)

func Run(gormDB *gorm.DB, direction Direction) error {
	return nil
}

// you have to refer to files - migrate and migration.go
// you have to do commands integration\
// config intgration using viper
