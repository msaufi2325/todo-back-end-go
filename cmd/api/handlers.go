package main

import (
	"fmt"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Todo API is up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.DB.AllTodos(1)
	if err != nil {
		fmt.Println(err)
	}

	_ = app.writeJSON(w, http.StatusOK, todos)
}
