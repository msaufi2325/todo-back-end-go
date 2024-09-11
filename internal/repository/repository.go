package repository

import (
	"database/sql"

	"github.com/msaufi2325/todo-back-end-go/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllTodos(id int) ([]*models.Todo, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}
