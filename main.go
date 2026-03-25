package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type application struct {
	db *sql.DB
}

func main() {
	dsn := "postgres://postgres:postgres@localhost:5432/university?sslmode=disable"

	db, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &application{db: db}

	mux := app.routes()

	log.Println("Server running on :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}