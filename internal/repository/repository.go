package repository

import "github.com/msaufi2325/todo-back-end-go/internal/models"

type DatabaseRepo interface {
	AllTodos(id int) ([]*models.Todo, error)
}
