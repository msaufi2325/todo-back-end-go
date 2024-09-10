package main

import (
	"errors"
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
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, todos)
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// read the json payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// validate the user against database
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid email or password"), http.StatusBadRequest)
		return
	}

	// check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid email or password"), http.StatusBadRequest)
		return
	}

	// create a jwt user
	u := jwtUser{
		ID:       user.ID,
		Username: user.UserName,
	}

	// generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// log.Println(tokens.Token)
	// go to jwt.io site to inspect the token

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	_ = app.writeJSON(w, http.StatusOK, tokens)
}
