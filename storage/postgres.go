package storage

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewPostgresConnection() (*gorm.DB, error) {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" {
		return nil, fmt.Errorf("incomplete database connection parameters")
	}

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		return db, err
	}

	return db, nil
}
