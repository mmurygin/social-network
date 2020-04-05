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

func Users() (users []User, err error) {
	rows, err := db.Query("SELECT id, name, surname, city FROM users")

	if err != nil {
		return
	}

	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Name, &user.Surname, &user.City); err != nil {
			return
		}

		users = append(users, user)
	}

	rows.Close()

	return
}
