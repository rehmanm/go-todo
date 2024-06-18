package data

import (
	"database/sql"
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

	query := `
			INSERT INTO todos (title)
			VALUES ($1)
			RETURNING id`

	args := []any{todo.Title}
	return m.DB.QueryRow(query, args...).Scan(&todo.ID)
}

func (m TodoModel) Get(id int64) (*Todo, error) {
	return nil, nil
}

func (m TodoModel) Update(todo *Todo) error {
	return nil
}

func (m TodoModel) Delete(todo *Todo) error {
	return nil
}
