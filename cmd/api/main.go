package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/msaufi2325/todo-back-end-go/internal/repository"
	"github.com/msaufi2325/todo-back-end-go/internal/repository/dbrepo"
)

const port = 8081

type application struct {
	DSN          string
	Domain       string
	DB           repository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
}

func main() {
		// Load .env file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

	// set application config
	var app application

	// Read JWT secret from environment variable
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
			fmt.Println("JWT_SECRET environment variable is not set")
			return
	}

	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5434 user=postgres password=password dbname=todos sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.StringVar(&app.JWTSecret, "jwt-secret", jwtSecret, "signing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "sigining issuer") // for development example.com is used
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "signing audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&app.Domain, "domain", "example.com", "domain")
	flag.Parse()

	// connect to database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.auth = Auth{
		Issuer: app.JWTIssuer,
		Audience: app.JWTAudience,
		Secret: app.JWTSecret,
		TokenExpiry: 15 * time.Minute,
		RefreshExpiry: 24 * time.Hour,
		CookiePath: "/",
		CookieName: "__Host-refresh_token",
		CookieDomain: app.CookieDomain,
	}

	log.Println("Starting server on", port)

	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
