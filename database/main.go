package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	urlConnection := "host=localhost port=7777 user=golang password=golang dbname=golang_db sslmode=disable"

	database, err := sql.Open("postgres", urlConnection)

	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	if err = database.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(database)
}
