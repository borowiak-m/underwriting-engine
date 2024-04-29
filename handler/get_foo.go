package handler

import "net/http"

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("from HandleFoo"))
	return nil
}
