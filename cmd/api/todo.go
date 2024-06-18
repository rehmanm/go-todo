package main

import (
	"fmt"
	"net/http"
	"time"

	"rehmanm.go-todo/internal/data"
)

func (app *application) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create todo")
}

func (app *application) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete todo")
}

func (app *application) getAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get all todos")
}

func (app *application) updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update todo")
}

func (app *application) getTodoHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	todo := data.Todo{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "First Todo",
		Completed: false,
		UpdatedAt: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"todo": todo}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
