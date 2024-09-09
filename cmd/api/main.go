package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8081

type application struct {
	DSN string
	Domain string
}

func main() {
	// set application config
	var app application

	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5434 user=postgres password=password dbname=todos sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connect to database

	app.Domain = "example.com"

	log.Println("Starting server on", port)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
