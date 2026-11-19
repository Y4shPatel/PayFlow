package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Y4shPatel/PayFlow/Internal/customer"
	"github.com/Y4shPatel/PayFlow/Internal/database"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	pool, err := database.Connect(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	log.Println("Successfully connected to PostgreSQL!")

	// Create router
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc(
		"GET /api/v1/customers/{id}",
		customer.GetCustomerHandler(pool),
	)

	// Start HTTP server
	log.Println("Server running on :8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
