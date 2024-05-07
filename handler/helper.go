package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ApiError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(statusCode int, err error) *ApiError {
	return &ApiError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

type APIfunc func(w http.ResponseWriter, r *http.Request) error

func Make(h APIfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiErr, ok := err.(*ApiError); ok {
				writeJSON(w, apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"statusCode": http.StatusInternalServerError,
					"message":    "internal server error",
				}
				writeJSON(w, http.StatusInternalServerError, errResp)
			}
			slog.Error("HTTP API error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
