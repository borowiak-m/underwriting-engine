package handler

import (
	"fmt"
	"net/http"

	"github.com/borowiak-m/underwriting-engine/data"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetCustomer(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return NewApiError(http.StatusBadRequest, err)
	}
	fmt.Println("fetching customer id:", id)
	customer := data.Customer{
		ID: id,
	}
	return writeJSON(w, http.StatusOK, customer)
}
