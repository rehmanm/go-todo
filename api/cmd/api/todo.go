package main

import (
	"fmt"
	"net/http"

	"rehmanm.go-todo/internal/data"
	"rehmanm.go-todo/internal/validator"
)

func (app *application) createTodoHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title string `json:"title"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	todo := &data.Todo{
		Title: input.Title,
	}

	v := validator.New()

	if data.ValidateTodo(v, todo); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Todos.Insert(todo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/todos/%d", todo.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"todo": todo}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
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

	todo, err := app.models.Todos.Get(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"todo": todo}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
