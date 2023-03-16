package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DatabaseConnection() *sql.DB {
	connection := "user=postgres dbname=go_store password=130103 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}
