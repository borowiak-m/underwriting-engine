package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/borowiak-m/underwriting-engine/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// config
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
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
