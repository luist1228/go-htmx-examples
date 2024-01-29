package handler

import "github.com/luist1228/go-htmx-examples/db"

type Handler struct {
	todos db.Todos
}

func NewHandler() *Handler {
	todos := db.FillTodos()
	return &Handler{
		todos: *todos,
	}
}
