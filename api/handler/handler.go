package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/luist1228/go-htmx-examples/db"
)

const (
	htmxHeaderKey = "HX-Request"
)

type Handler struct {
	todos db.Todos
}

func NewHandler() *Handler {
	todos := db.FillTodos()
	return &Handler{
		todos: *todos,
	}
}

func (h Handler) isHtmx(c fiber.Ctx) bool {
	return c.Get(htmxHeaderKey) == "true"
}
