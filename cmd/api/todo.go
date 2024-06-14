package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of the todo %d\n", id)
}
