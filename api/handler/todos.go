package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/luist1228/go-htmx-examples/templates/components"
	"github.com/luist1228/go-htmx-examples/templates/views"
)

func (h *Handler) RegisterTodosRequests(app *fiber.App, api fiber.Router) {
	// Views
	app.Get("/todos", h.getTodosRequest)
	app.Post("/todos", h.addTodoRequest)

	// API
	api.Get("/todos", h.getTodosRequest)
	api.Post("/todos", h.addTodoRequest)

}

func (h *Handler) getTodosRequest(c fiber.Ctx) error {
	if isApiRequest(c) {
		return c.JSON(h.todos.All())
	}

	content := views.TodosView(h.todos)
	if isHtmx(c) {
		return Render(c, content)
	}

	return Render(c, FullPageRender("TODOS", content))
}

type addTodoRequest struct {
	Description string `form:"description" validate:"required"`
}

func (h *Handler) addTodoRequest(c fiber.Ctx) error {
	var body addTodoRequest
	if err := c.Bind().Form(&body); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	if err := h.validate.Struct(body); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	todo := h.todos.Add(body.Description, false)

	if isApiRequest(c) {
		return c.JSON(todo)
	}

	if isHtmx(c) {
		return Render(c, components.Todo(*todo))
	}

	return Render(c, FullPageRender("TODOS", views.TodosView(h.todos)))
}
