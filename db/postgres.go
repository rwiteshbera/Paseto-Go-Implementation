package db

import (
	"github.com/rwiteshbera/Paseto-Go-Implementation/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config config.Config) (*gorm.DB, error) {
	database, err := gorm.Open(postgres.Open(config.DB_CONNECTION_STRING), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return database, nil
}
