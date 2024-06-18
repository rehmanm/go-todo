package data

import (
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
