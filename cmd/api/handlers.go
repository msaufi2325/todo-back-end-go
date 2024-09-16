package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/msaufi2325/todo-back-end-go/internal/models"
	"golang.org/x/crypto/bcrypt"
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
	var requestPayload struct {
		UserID int `json:"user_id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	todos, err := app.DB.AllTodos(requestPayload.UserID)
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
		ID:       strconv.Itoa(user.ID),
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

	data := map[string]string{
		"access_token":  tokens.Token,
		"refresh_token": tokens.RefreshToken,
		"username":      user.UserName,
		"user_id":       strconv.Itoa(user.ID),
	}

	_ = app.writeJSON(w, http.StatusOK, data)
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestPayload.Password), 12)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// create a user
	user := models.User{
		UserName:  requestPayload.Username,
		Email:     requestPayload.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	// insert the user into the database
	id, err := app.DB.InsertUser(user)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// create a jwt user
	u := jwtUser{
		ID:       strconv.Itoa(id),
		Username: user.UserName,
	}

	// generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// set refresh token as a cookie
	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	resp := JSONResponse{
		Error:   false,
		Message: "User created successfully",
		Data: map[string]string{
			"access_token":  tokens.Token,
			"refresh_token": tokens.RefreshToken,
			"user_id":       u.ID,
			"username":      u.Username,
		},
	}

	_ = app.writeJSON(w, http.StatusCreated, resp)
}

func (app *application) refreshToken(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value
			// parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(app.JWTSecret), nil
			})
			if err != nil {
				log.Println("Error parsing token:", err)
				app.errorJSON(w, errors.New("unauthorized: invalid token"), http.StatusUnauthorized)
				return
			}

			// get the user id from the token claims
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				log.Println("Error converting user ID:", err)
				app.errorJSON(w, errors.New("unauthorized: invalid user ID"), http.StatusUnauthorized)
				return
			}

			// get the user from the database
			user, err := app.DB.GetUserByID(userID)
			if err != nil {
				log.Println("Error fetching user from DB:", err)
				app.errorJSON(w, errors.New("unauthorized: user not found"), http.StatusUnauthorized)
				return
			}

			u := jwtUser{
				ID:       strconv.Itoa(user.ID),
				Username: user.UserName,
			}

			tokenPairs, err := app.auth.GenerateTokenPair(&u)
			if err != nil {
				app.errorJSON(w, errors.New("error generating tokens"), http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, app.auth.GetRefreshCookie(tokenPairs.RefreshToken))

			app.writeJSON(w, http.StatusOK, tokenPairs)
		}
	}
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
	app.writeJSON(w, http.StatusOK, "Logged out")
}

func (app *application) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var requestPayload models.Todo

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	todo, err := app.DB.OneTodo(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	todo.Title = requestPayload.Title
	todo.Description = requestPayload.Description
	todo.Category = requestPayload.Category
	todo.Priority = requestPayload.Priority
	todo.IsCompleted = requestPayload.IsCompleted
	todo.IsRemoved = requestPayload.IsRemoved
	todo.UpdatedAt = time.Now().UTC()

	err = app.DB.UpdateTodo(*todo)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "Todo updated successfully",
	}

	_ = app.writeJSON(w, http.StatusOK, resp)
}

func (app *application) AddTodo(w http.ResponseWriter, r *http.Request) {
	var requestPayload models.Todo

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	requestPayload.CreatedAt = time.Now().UTC()
	requestPayload.UpdatedAt = time.Now().UTC()
	requestPayload.ID = 0 // Ensure ID is zero to avoid conflicts

	newID, err := app.DB.InsertTodo(requestPayload)
	if err != nil {
		log.Println("Error inserting todo:", err)
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "Todo added successfully",
		Data: map[string]int{
			"id":      newID,
			"user_id": requestPayload.UserID,
		},
	}

	_ = app.writeJSON(w, http.StatusCreated, resp)
}

func (app *application) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.DB.DeleteTodoByID(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "Todo deleted successfully",
	}

	_ = app.writeJSON(w, http.StatusOK, resp)
}
