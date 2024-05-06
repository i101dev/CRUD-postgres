package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/i101dev/rss-aggregator/controllers"
)

type Database struct {
	DB *gorm.DB
}

var (
	db *gorm.DB
)

func GetDB() *gorm.DB {
	return db
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

	d, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		return db, err
	} else {
		db = d
	}

	controllers.InitUsers(d)

	return d, nil
}
