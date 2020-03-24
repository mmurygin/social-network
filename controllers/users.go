package controllers

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/mmurygin/social-network/data"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user := new(data.User)
	decoder := schema.NewDecoder()

	err = decoder.Decode(user, r.PostForm)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = user.Create()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Printf("%+v\n", user)

	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
