package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/msaufi2325/todo-back-end-go/internal/repository"
	"github.com/msaufi2325/todo-back-end-go/internal/repository/dbrepo"
)

const port = 8081

type application struct {
	DSN string
	Domain string
	DB repository.DatabaseRepo
}

func main() {
	// set application config
	var app application

	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5434 user=postgres password=password dbname=todos sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connect to database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.Domain = "example.com"

	log.Println("Starting server on", port)

	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
