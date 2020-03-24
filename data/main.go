package data

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db sql.DB

func init() {
	connectionStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
}
