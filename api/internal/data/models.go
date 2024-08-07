package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Todos interface {
		Insert(todo *Todo) error
		Get(id int64) (*Todo, error)
		Update(todo *Todo) error
		Delete(id int64) error
	}
	Users interface {
		Insert(user *User) error
		GetByEmail(email string) (*User, error)
		Update(todo *User) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Todos: TodoModel{DB: db},
		Users: UserModel{DB: db},
	}
}
