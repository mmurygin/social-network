package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if token != "" {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func main() {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	r.Use(authMiddleware)
	r.HandleFunc("/", indexHandler)
	r.NotFoundHandler = handlers.LoggingHandler(os.Stdout, http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8000", r))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Index")
	fmt.Fprintf(w, "Index\n")
}
