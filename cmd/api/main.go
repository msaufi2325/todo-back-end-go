package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8081

type application struct {
	Domain string
}

func main() {
	// set application config
	var app application

	// read from command line

	// connect to database

	app.Domain = "example.com"

	log.Println("Starting server on", port)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}