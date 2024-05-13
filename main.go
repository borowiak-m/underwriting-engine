package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/borowiak-m/underwriting-engine/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func getDB() *bun.DB {
	var (
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		name     = os.Getenv("DB_NAME")
	)

	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		user,
		password,
		host,
		name)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	return bun.NewDB(sqldb, pgdialect.New())
}

func main() {
	// config
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	// db
	db := getDB()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// routes
	router := chi.NewMux()
	router.Get("/customer/{id}", handler.Make(handler.GetCustomer))
	router.Post("/upload", handler.Make(handler.Upload))
	router.Post("/file", handler.Make(handler.CreateFileUpload))

	// server
	PORT := os.Getenv("PORT")
	slog.Info("API server running", "address", PORT)
	http.ListenAndServe(PORT, router)
}
