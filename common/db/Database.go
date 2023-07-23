package db

import (
	"fmt"
	"gotest/common/envs"
	"gotest/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(config *envs.DbConfig) (*Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Host, config.User, config.Password, config.Name, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}

func (db *Database) InitScheme() error {
	return db.DB.AutoMigrate(&models.UserDto{})
}
