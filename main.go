package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/i101dev/rss-aggregator/models"
	"github.com/i101dev/rss-aggregator/routes"
	"github.com/i101dev/rss-aggregator/storage"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type Database struct {
	DB *gorm.DB
}

func main() {

	fmt.Println("Starting up...")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Invalid port - not found in environment")
	}

	// -----------------------------------------------------------------------
	// Database Setup
	//

	db, err := storage.NewPostgresConnection()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	} else {
		fmt.Println("Postgres connection success!")
	}

	if err = models.MigrateUsers(db); err != nil {
		log.Fatal("could not migrate [users]")
	}

	// -----------------------------------------------------------------------
	// Server Setup
	//

	router := routes.BuildRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Println("Server is live on port:", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
