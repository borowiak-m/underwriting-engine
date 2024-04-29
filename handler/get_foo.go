package handler

import "net/http"

type Response struct {
	Content string
}

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	userName := Response{
		Content: "YourName",
	}
	return writeJSON(w, http.StatusOK, userName)
}
