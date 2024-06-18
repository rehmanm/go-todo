package data

import "time"

type Todo struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // to hide the value
	Title     string    `json:"title"`
	Completed bool      `json:"completed"` // omitemtpy -> to hide the value if false
	UpdatedAt time.Time `json:"-"`
}

//string directive can be used if user wants to convert numerical values to string
