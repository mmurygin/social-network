package controllers

import (
	"github.com/gorilla/schema"
	"github.com/mmurygin/social-network/data"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	user := new(data.User)
	decoder := schema.NewDecoder()
	log.Println(r.PostForm)

	err = decoder.Decode(user, r.PostForm)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	log.Println(user)
}
