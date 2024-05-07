package main

import (
	"log/slog"
	"net/http"

	"github.com/borowiak-m/underwriting-engine/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	var envConfig map[string]string

	envConfig, err := godotenv.Read()
	if err != nil {
		panic("Error loading .env file")
	}
	PORT := envConfig["PORT"]
	router := chi.NewMux()
	router.Get("/customer/{id}", handler.Make(handler.GetCustomer))
	slog.Info("API server running", "address", PORT)
	http.ListenAndServe(PORT, router)
}
