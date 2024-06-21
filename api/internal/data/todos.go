package data

import (
	"database/sql"
	"errors"
	"fmt"
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
	return m.DB.QueryRow(query, args...).Scan(&todo.ID)
}

func (m TodoModel) Get(id int64) (*Todo, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT * from todo_get($1)`

	var todo Todo

	err := m.DB.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &todo, nil
}

func (m TodoModel) Update(todo *Todo) error {
	return nil
}

func (m TodoModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `call todo_delete( $1)`

	var rowsAffected int64
	err := m.DB.QueryRow(query, id).Scan(&rowsAffected)
	if err != nil {
		return err
	}

	fmt.Printf("rowsAffected %d", rowsAffected)

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
