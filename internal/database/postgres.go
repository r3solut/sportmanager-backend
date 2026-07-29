package database

import (
	"fmt"
	"log"
	"://github.com"
	"://github.com"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil { return nil, fmt.Errorf("failed to connect to database: %w", err) }
	if err := db.AutoMigrate(&domain.User{}, &domain.Wallet{}); err != nil { return nil, fmt.Errorf("failed to auto migrate tables: %w", err) }
	log.Println("Database connection established successfully")
	return db, nil
}
