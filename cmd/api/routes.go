package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Get("/", app.Home)

	mux.Post("/authenticate", app.authenticate)
	mux.Post("/register", app.register)
	mux.Get("/refresh", app.refreshToken)
	mux.Get("/logout", app.logout)

	mux.Route("/todos", func(mux chi.Router) {
		mux.Use(app.authRequired)

		mux.Get("/all", app.AllTodos)
		mux.Put("/0", app.AddTodo)
		mux.Patch("/{id}", app.UpdateTodo)
		mux.Delete("/{id}", app.DeleteTodo)
	})

	return mux
}
