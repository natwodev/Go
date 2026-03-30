package postgres

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection() (*gorm.DB, error) {
	host := getenv("DB_HOST", "localhost")
	user := getenv("DB_USER", "postgres")
	password := getenv("DB_PASSWORD", "postgres")
	dbname := getenv("DB_NAME", "go_backend")
	port := getenv("DB_PORT", "5432")
	sslmode := getenv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	log.Println("database connection established")
	return db, nil
}

func getenv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
