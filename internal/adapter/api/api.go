package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/avazquezcode/govetryx/internal/adapter/interpreter"
)

type interpretRequest struct {
	SourceCode string
}

func InterpretHandler(w http.ResponseWriter, r *http.Request) {
	var request interpretRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		jsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if request.SourceCode == "" {
		jsonResponse(w, "source code cannot be empty", http.StatusBadRequest)
		return
	}

	output, err := interpreter.RunCode(request.SourceCode)
	if err != nil {
		jsonResponse(w, fmt.Sprintf("error on the interpreter: %s", err.Error()), http.StatusBadRequest)
		return
	}

	jsonResponse(w, output, http.StatusAccepted)
}
