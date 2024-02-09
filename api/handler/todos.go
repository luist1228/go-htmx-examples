package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/luist1228/go-htmx-examples/templates/views"
)

func (h *Handler) RegisterTodosRequests(app *fiber.App, api fiber.Router) {
	// Views
	app.Get("/todos", h.getTodosRequestView)

	// API
	api.Get("/todos", h.getTodosRequest)
}

func (h *Handler) getTodosRequestView(c fiber.Ctx) error {
	content := views.TodosView(h.todos)
	if h.isHtmx(c) {
		return Render(c, content)
	}

	return Render(c, FullPageRender("TODOS", content))
}

func (h *Handler) getTodosRequest(c fiber.Ctx) error {
	return c.JSON(h.todos.All())
}
