package handler

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/luist1228/go-htmx-examples/templates/components"
	"github.com/luist1228/go-htmx-examples/templates/views"
)

func (h *Handler) RegisterTodosRequests(app *fiber.App, api fiber.Router) {
	// Views
	app.Get("/todos", h.getTodosRequest)
	app.Post("/todos", h.addTodoRequest)
	app.Post("/todos/delete/:id", h.addTodoRequest)
	app.Delete("/todos/:id", h.deleteTodoRequest)

	// API
	api.Get("/todos", h.getTodosRequest)
	api.Post("/todos", h.addTodoRequest)
	api.Delete("/todos/:id", h.deleteTodoRequest)

}

func (h Handler) fullPageTodoComponent() templ.Component {
	return FullPageRender("TODOS", views.TodosView(h.todos))
}

func (h *Handler) getTodosRequest(c fiber.Ctx) error {
	content := views.TodosView(h.todos)

	return h.caseResponse(
		c,
		content,
		h.todos.All(),
		h.fullPageTodoComponent(),
	)
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

	return h.caseResponse(
		c,
		components.Todo(*todo),
		c.JSON(todo),
		h.fullPageTodoComponent(),
	)

}

type deleteTodoRequest struct {
	ID string `uri:"id" validate:"required,uuid"`
}

func (h *Handler) deleteTodoRequest(c fiber.Ctx) error {
	var params deleteTodoRequest

	if err := c.Bind().URI(&params); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	if err := h.validate.Struct(params); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}
	fmt.Println(h.todos.All())
	todoID, _ := uuid.Parse(params.ID)
	h.todos.Remove(todoID)
	fmt.Println(h.todos.All())
	return h.caseResponse(
		c,
		components.Todos(h.todos),
		h.todos,
		h.fullPageTodoComponent(),
	)
}
