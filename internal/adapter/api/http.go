package api

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string `json:"message"`
}

func jsonResponse(w http.ResponseWriter, message string, statusCode int) {
	p := response{
		Message: message,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(p)
}
