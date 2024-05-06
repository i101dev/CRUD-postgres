package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/i101dev/rss-aggregator/config"
	"github.com/i101dev/rss-aggregator/routes"
	"github.com/joho/godotenv"
)

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

	_, err := config.NewPostgresConnection()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	} else {
		fmt.Println("Postgres connection success!")
	}

	// -----------------------------------------------------------------------
	// Server Setup
	//

	router := routes.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Println("Server is live on port:", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
