// Package main provides the entrypoint for the Budget API service.
// It wires configuration, database connectivity, and HTTP routing.
package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/yadieloscar/budget-api/internal/api"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// main bootstraps the application by loading environment variables,
	// opening a PostgreSQL connection, building the Gin router, and starting
	// the HTTP server. See README for required env vars.
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		// Example:
		// export DB_DSN='postgres://user:pass@localhost:5432/budget?sslmode=disable'
		log.Fatal("DB_DSN is not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	r := api.SetupRouter(db)

	r.Run() // listen and serve on 0.0.0.0:8080
}
