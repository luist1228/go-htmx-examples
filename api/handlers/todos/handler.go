package todos

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/luist1228/go-htmx-examples/api/handlers"
	"github.com/luist1228/go-htmx-examples/db"
	"github.com/luist1228/go-htmx-examples/templates/components"
	"github.com/luist1228/go-htmx-examples/templates/views"
)

type (
	Handler interface {
		GetTodos(c fiber.Ctx) error
		Mount(app *fiber.App, api fiber.Router)
	}

	handler struct {
		todos    db.Todos
		validate *validator.Validate
	}
)

func NewHandler(validator *validator.Validate) Handler {
	todos := db.FillTodos()
	return &handler{
		todos:    *todos,
		validate: validator,
	}
}

func (h *handler) Mount(app *fiber.App, api fiber.Router) {
	// Views
	app.Get("/todos", h.GetTodos)
	app.Post("/todos", h.AddTodo)
	// Native Post to delete
	app.Post("/todos/delete/:id", h.DeleteTodo)
	// With Htmx delete
	app.Delete("/todos/:id", h.DeleteTodo)

	// API
	api.Get("/todos", h.GetTodos)
	api.Post("/todos", h.AddTodo)
	api.Delete("/todos/:id", h.DeleteTodo)
}

func (h *handler) GetTodos(c fiber.Ctx) error {
	content := views.TodosView(h.todos)

	return handlers.CaseResponse(
		c,
		content,
		h.todos.All(),
		h.fullPageTodoComponent(),
	)
}

type addTodoRequest struct {
	Description string `form:"description" validate:"required"`
}

func (h *handler) AddTodo(c fiber.Ctx) error {
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

	return handlers.CaseResponse(
		c,
		components.Todo(*todo),
		c.JSON(todo),
		h.fullPageTodoComponent(),
	)
}

type deleteTodo struct {
	ID string `uri:"id" validate:"required,uuid"`
}

func (h *handler) DeleteTodo(c fiber.Ctx) error {
	var params deleteTodo

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
	return handlers.CaseResponse(
		c,
		components.Todos(h.todos),
		h.todos,
		h.fullPageTodoComponent(),
	)
}

func (h *handler) fullPageTodoComponent() templ.Component {
	return handlers.FullPageRender("TODOS", views.TodosView(h.todos))
}
