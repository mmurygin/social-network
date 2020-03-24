package data

import (
	_ "github.com/gorilla/schema"
)

type User struct {
	Id        int    `schema:"-"`
	Email     string `schema:"email"`
	Password  string `schema:"password"`
	Name      string `schema:"name"`
	Surname   string `schema:"surname"`
	Age       int    `schema:"age"`
	Gender    string `schema:"gender"`
	Interests string `schema:"interests"`
	City      string `schema:"city"`
}
