package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB, error) {
	urlConnection := "host=localhost port=7777 user=postgres password=admin dbname=golang_db sslmode=disable"

	database, err := sql.Open("postgres", urlConnection)

	if err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		return nil, err
	}
	return database, nil
}
