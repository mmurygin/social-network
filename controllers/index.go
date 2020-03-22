package controllers

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	serveTemplate(w, r, "index.tmpl", nil)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	serveTemplate(w, r, "signup.tmpl", nil)
}
