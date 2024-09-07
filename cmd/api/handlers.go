package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/msaufi2325/todo-back-end-go/internal/models"
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

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllTodos(w http.ResponseWriter, r *http.Request) {
	var DummyTodoList = []models.Todo{
		{
			ID:          1,
			Title:       "Work on the project 1",
			Description: "Work on the project for 2 hours",
			Category:    "work",
			Priority:    "high",
			IsCompleted: false,
			IsRemoved:   false,
			CreatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UserID:      1,
		},
		{
			ID:          2,
			Title:       "Buy groceries",
			Description: "Buy groceries for the week",
			Category:    "home",
			Priority:    "medium",
			IsCompleted: false,
			IsRemoved:   false,
			CreatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UserID:      1,
		},
		{
			ID:          3,
			Title:       "Workout",
			Description: "Workout for 1 hour",
			Category:    "hobby",
			Priority:    "low",
			IsCompleted: false,
			IsRemoved:   false,
			CreatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UserID:      1,
		},
		{
			ID:          4,
			Title:       "Work on the project 2",
			Description: "Work on the project for 2 hours",
			Category:    "work",
			Priority:    "high",
			IsCompleted: false,
			IsRemoved:   false,
			CreatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UserID:      1,
		},
		{
			ID:          5,
			Title:       "Buy groceries",
			Description: "Buy groceries for the week",
			Category:    "home",
			Priority:    "medium",
			IsCompleted: false,
			IsRemoved:   false,
			CreatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UserID:      1,
		},
		{
			ID:          6,
			Title:       "Workout",
			Description: "Workout for 1 hour",
			Category:    "hobby",
			Priority:    "low",
			IsCompleted: false,
			IsRemoved:   false,
			CreatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UserID:      1,
		},
		{
			ID:          7,
			Title:       "Work on the project 3",
			Description: "Work on the project for 2 hours",
			Category:    "work",
			Priority:    "high",
			IsCompleted: false,
			IsRemoved:   false,
			CreatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UserID:      1,
		},
		{
			ID:          8,
			Title:       "Buy groceries 2",
			Description: "Buy groceries for the week",
			Category:    "home",
			Priority:    "medium",
			IsCompleted: true,
			IsRemoved:   true,
			CreatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UserID:      1,
		},
		{
			ID:          9,
			Title:       "Work on the project 4",
			Description: "Work on the project for 2 hours",
			Category:    "work",
			Priority:    "high",
			IsCompleted: false,
			IsRemoved:   false,
			CreatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UserID:      1,
		},
		{
			ID:          10,
			Title:       "Buy groceries 3",
			Description: "Buy groceries for the week",
			Category:    "home",
			Priority:    "medium",
			IsCompleted: false,
			IsRemoved:   false,
			CreatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UserID:      1,
		},
		{
			ID:          11,
			Title:       "Workout 2",
			Description: "Workout for 1 hour",
			Category:    "hobby",
			Priority:    "low",
			IsCompleted: false,
			IsRemoved:   false,
			CreatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UserID:      1,
		},
		{
			ID:          12,
			Title:       "Work on the project 5",
			Description: "Work on the project for 2 hours",
			Category:    "work",
			Priority:    "high",
			IsCompleted: true,
			IsRemoved:   true,
			CreatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UserID:      1,
		},
		{
			ID:          13,
			Title:       "Work on the project 5",
			Description: "Work on the project for 2 hours",
			Category:    "work",
			Priority:    "high",
			IsCompleted: true,
			IsRemoved:   true,
			CreatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
			UserID:      1,
		},
	}

	out, err := json.Marshal(DummyTodoList)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
