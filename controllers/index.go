package controllers

import (
	"github.com/mmurygin/social-network/auth"
	"github.com/mmurygin/social-network/data"
	"log"
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
	if r.Method == "GET" {
		serveTemplate(w, r, "signup.html", nil)
	} else {
		log.Println("Start auth")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		userId, err := data.CheckAndQueryUser(email, password)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		err = auth.StoreSession(w, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		serveTemplate(w, r, "signin.html", nil)
		return
	}

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	userId, err := data.CheckAndQueryUser(email, password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = auth.StoreSession(w, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	auth.CleanSession(w)

	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}
