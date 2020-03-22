package controllers

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	serveTemplate(w, r, "index.tmpl", nil)
}
