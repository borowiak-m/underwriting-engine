package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type APIfunc func(w http.ResponseWriter, r *http.Request) error

func Make(h APIfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("HTTP API error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
