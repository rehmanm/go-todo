package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"rehmanm.go-todo/internal/validator"
)

type Todo struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // to hide the value
	Title     string    `json:"title"`
	Completed bool      `json:"completed"` // omitemtpy -> to hide the value if false
	UpdatedAt time.Time `json:"-"`
}

func ValidateTodo(v *validator.Validator, t *Todo) {
	v.Check(t.Title != "", "title", "must be provided")
	v.Check(len(t.Title) <= 500, "title", "must not be more than 500 bytes long")
}

//string directive can be used if user wants to convert numerical values to string

type TodoModel struct {
	DB *sql.DB
}

func (m TodoModel) Insert(todo *Todo) error {

	query := `call todo_save(0, $1, false)`

	args := []any{todo.Title}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&todo.ID)
}

func (m TodoModel) Get(id int64) (*Todo, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT  * from todo_get($1)`

	var todo Todo

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &todo, nil
}

func (m TodoModel) Update(todo *Todo) error {

	query := `
		UPDATE todos
		SET title = $1, completed = $2, updated_at = CURRENT_TIMESTAMP
		WHERE id = $3
		RETURNING id
	`

	args := []any{
		todo.Title,
		todo.Completed,
		todo.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&todo.ID)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil

}

func (m TodoModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `call todo_delete( $1)`

	var rowsAffected int64

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&rowsAffected)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
