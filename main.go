package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mmurygin/social-network/controllers"
	_ "github.com/mmurygin/social-network/data"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if token != "" {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

func notFoundHandler() http.Handler {
	return handlers.LoggingHandler(os.Stdout, http.NotFoundHandler())
}

func main() {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	r.NotFoundHandler = notFoundHandler()

	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	r.HandleFunc("/", controllers.Index)
	r.HandleFunc("/signup", controllers.SignUp)
	r.HandleFunc("/signin", controllers.SignIn)
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", controllers.ViewUser)

	log.Fatal(http.ListenAndServe(":3001", r))
}
