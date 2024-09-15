package dbrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/msaufi2325/todo-back-end-go/internal/models"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AllTodos(id int) ([]*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select
			t.id, t.title, coalesce(t.description, ''), 
			t.category, t.priority, t.is_completed, t.is_removed, 
			t.created_at, t.updated_at, t.user_id
		from
			todos t
		join 
			users u on t.user_id = u.id
		where
			u.id = $1
		order by
			id asc
	`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*models.Todo

	for rows.Next() {
		var t models.Todo
		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Category,
			&t.Priority,
			&t.IsCompleted,
			&t.IsRemoved,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.UserID,
		)
		if err != nil {
			return nil, err
		}

		todos = append(todos, &t)
	}

	return todos, nil
}

func (m *PostgresDBRepo) OneTodo(id int) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select
			t.id, t.title, coalesce(t.description, ''),
			t.category, t.priority, t.is_completed, t.is_removed,
			t.created_at, t.updated_at, t.user_id
		from
			todos t
		join
			users u on t.user_id = u.id
		where
			t.id = $1
	`
	
	row := m.DB.QueryRowContext(ctx, query, id)

	var todo models.Todo

	err := row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Category,
		&todo.Priority,
		&todo.IsCompleted,
		&todo.IsRemoved,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.UserID,
	)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (m *PostgresDBRepo) UpdateTodo(todo models.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		update
			todos
		set
			title = $1,
			description = $2,
			category = $3,
			priority = $4,
			is_completed = $5,
			is_removed = $6,
			updated_at = $7
		where
			id = $8
	`

	_, err := m.DB.ExecContext(
		ctx,
		query,
		todo.Title,
		todo.Description,
		todo.Category,
		todo.Priority,
		todo.IsCompleted,
		todo.IsRemoved,
		time.Now(),
		todo.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select
			id, username, email, password, created_at, updated_at
		from
			users
		where
			email = $1
	`

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *PostgresDBRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select
			id, username, email, password, created_at, updated_at
		from
			users
		where
			id = $1
	`

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
