package database

import (
	"SubscriptionService/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func InitDB(cfg *config.DBConfig) (*Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}

func (db *Database) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
