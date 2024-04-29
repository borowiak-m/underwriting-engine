package handler

import (
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
