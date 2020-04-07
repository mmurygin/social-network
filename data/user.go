package data

import (
	_ "github.com/gorilla/schema"
	"golang.org/x/crypto/bcrypt"
	"log"
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
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

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
	rows, err := db.Query(
		"SELECT id, name, surname, city FROM users order by createdAt desc")

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

func QueryUser(id int) (*User, error) {
	log.Println("Query User")

	row := db.QueryRow(`
		SELECT id, email, name, surname, age, gender, interests, city
		FROM users
		WHERE id = ?;`, id)

	user := User{}
	err := row.Scan(&user.Id, &user.Email, &user.Name, &user.Surname, &user.Age,
		&user.Gender, &user.Interests, &user.City)

	return &user, err
}

func CheckAndQueryUser(email string, password string) (int, error) {
	log.Println("Query User")

	row := db.QueryRow(`
		SELECT id, password
		FROM users
		WHERE email = ?;`, email)

	user := User{}
	log.Println("Before scan")
	err := row.Scan(&user.Id, &user.Password)

	if err != nil {
		log.Println("Scan error")
		log.Println(err)
		return 0, err
	}

	log.Println("Compare passwords")
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 0, err
	}

	return user.Id, nil
}
