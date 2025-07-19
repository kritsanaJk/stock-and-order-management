package database

import (
	"stock-and-order-management/infrastructure/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB gorm.DB

func InitDatabase() error {

	db, err := gorm.Open(sqlite.Open(config.AppConfig.DatabaseLocation), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = *db

	return nil
}
