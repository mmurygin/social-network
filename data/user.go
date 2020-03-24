package data

import (
	_ "github.com/gorilla/schema"
	"time"
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
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) Create() error {
	statement := `
		INSERT INTO
		users (email, password, name, surname, age, gender, interests, city)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	res, err := db.Exec(statement, user.Email, user.Password, user.Name, user.Surname,
		user.Age, user.Gender, user.Interests, user.City)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	user.Id = int(id)

	return err
}
