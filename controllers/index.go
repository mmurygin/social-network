package controllers

import (
	"github.com/mmurygin/social-network/data"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	users, err := data.Users()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serveTemplate(w, r, "index.html", users)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	serveTemplate(w, r, "signup.html", nil)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	serveTemplate(w, r, "signin.html", nil)
}
