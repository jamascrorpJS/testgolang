package database

import (
	"fmt"

	"github.com/jamascrorpJS/eBank/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ConfigDataBase struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func NewClientDB(config ConfigDataBase) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DBName, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	db.Logger.LogMode(logger.Info)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Automigrate(config ConfigDataBase) error {
	db, err := NewClientDB(config)
	if err != nil {
		return err
	}
	err = db.AutoMigrate(
		&domain.Operation{},
		&domain.User{},
	)
	if err != nil {
		return err
	}
	return nil
}
