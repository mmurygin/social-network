package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mmurygin/social-network/auth"
	"github.com/mmurygin/social-network/controllers"
	"github.com/mmurygin/social-network/data"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := auth.GetSession(r)

		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		user, err := data.QueryUser(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		newContext := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(newContext)

		next.ServeHTTP(w, r)
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
	r.HandleFunc("/signup", controllers.SignUp)
	r.HandleFunc("/signin", controllers.SignIn).Methods("GET", "POST")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	secRoutes := r.PathPrefix("/").Subrouter()
	secRoutes.Use(authMiddleware)

	secRoutes.HandleFunc("/", controllers.Index)
	secRoutes.HandleFunc("/users/{id:[0-9]+}", controllers.ViewUser)

	log.Fatal(http.ListenAndServe(":3001", r))
}
