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

func authMiddleware(next http.Handler) http.Handler {
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

	apiRouter := r.PathPrefix("/api").Subrouter()
	initApiRouter(apiRouter)

	log.Fatal(http.ListenAndServe(":3001", r))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is test handler"))
}

func initApiRouter(apiRouter *mux.Router) {
	apiRouter.Use(authMiddleware)
	apiRouter.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	apiRouter.HandleFunc("/test", testHandler)
}
