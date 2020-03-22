package controllers

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	serveTemplate(w, r, "index.html", nil)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	serveTemplate(w, r, "signup.html", nil)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	serveTemplate(w, r, "signin.html", nil)
}
