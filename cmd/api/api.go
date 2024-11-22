package main

import (
	"fmt"
	"net/http"
	"os"

	api "github.com/avazquezcode/govetryx/internal/adapter/api"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to GoVetryx API!")
	})

	router.HandleFunc("/interpret", api.InterpretHandler).Methods("POST")
	router.Use(contentTypeMiddleware)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("ALLOWED_ORIGIN")},
	})

	handler := c.Handler(router)
	http.ListenAndServe(":8080", handler)
}

func contentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
