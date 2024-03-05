package todos

import (
	"encoding/json"
	"fmt"

	"github.com/a-h/templ"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/utils/v2"
	"github.com/google/uuid"
	"github.com/luist1228/go-htmx-examples/api/handlers"
	"github.com/luist1228/go-htmx-examples/db"
	"github.com/luist1228/go-htmx-examples/templates/components"
	"github.com/luist1228/go-htmx-examples/templates/views"
	"github.com/luist1228/go-htmx-examples/util"
)

type (
	Handler interface {
		GetTodos(c fiber.Ctx) error
		AddTodo(c fiber.Ctx) error
		Mount(app *fiber.App, api fiber.Router)
		SortTodos(c fiber.Ctx) error
		DeleteTodo(c fiber.Ctx) error
		GetEditTodo(c fiber.Ctx) error
		EditTodo(c fiber.Ctx) error
		SearchTodo(c fiber.Ctx) error
	}

	handler struct {
		validate *validator.Validate
	}
)

var todos = db.List

func NewHandler(validator *validator.Validate) Handler {
	fmt.Println("New Todo handler")
	return &handler{
		validate: validator,
	}
}

func (h *handler) Mount(app *fiber.App, api fiber.Router) {
	fmt.Println("Todo routes Mount")
	todosRouter := app.Group("/todos")
	// Views
	todosRouter.Get("/", h.GetTodos)
	todosRouter.Post("/", h.AddTodo)
	// Native Post to delete
	todosRouter.Post("/delete/:id", h.DeleteTodo)
	// With Htmx delete
	todosRouter.Delete("/:id", h.DeleteTodo)
	todosRouter.Post("/sort", h.SortTodos)

	todosRouter.Get("/edit/:id", h.GetEditTodo)
	todosRouter.Patch("/edit/:id", h.EditTodo)
	todosRouter.Post("/edit/:id", h.EditTodo)

	todosRouter.Get("/search", h.SearchTodo)

	// API
	api.Get("/todos", h.GetTodos)
	api.Post("/todos", h.AddTodo)
	api.Post("/todos/sort", h.SortTodos)
	api.Delete("/todos/:id", h.DeleteTodo)
}

func (h *handler) GetTodos(c fiber.Ctx) error {
	content := views.TodosView(*todos)

	return handlers.CaseResponse(
		c,
		content,
		todos,
		h.fullPageTodoComponent(nil),
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
	result := utils.CopyString(body.Description)
	todo := todos.Add(result)

	return handlers.CaseResponse(
		c,
		components.Todos(*todos),
		todo,
		h.fullPageTodoComponent(nil),
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
	todoID, _ := uuid.Parse(params.ID)
	todos.Remove(todoID)
	return handlers.CaseResponse(
		c,
		components.Todos(*todos),
		todos,
		h.fullPageTodoComponent(nil),
	)
}

func (h *handler) fullPageTodoComponent(t db.Todos) templ.Component {
	content := views.TodosView(*todos)
	if t != nil {
		content = views.TodosView(t)
	}
	return handlers.FullPageRender("TODOS", content)
}

type sortTodosRequest struct {
	Ids []string `form:"id"`
}

func (h *handler) SortTodos(c fiber.Ctx) error {
	var body sortTodosRequest = sortTodosRequest{
		Ids: make([]string, 0),
	}

	if err := c.Bind().MultipartForm(&body); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	ids := make([]uuid.UUID, 0)
	for _, id := range body.Ids {
		parsedId, _ := uuid.Parse(id)
		ids = append(ids, parsedId)
	}
	todos.Reorder(ids)

	return handlers.CaseResponse(
		c,
		components.Todos(*todos),
		*todos,
		h.fullPageTodoComponent(nil),
	)

}

type getEditTodo struct {
	ID string `uri:"id" validate:"required,uuid"`
}

func (h *handler) GetEditTodo(c fiber.Ctx) error {
	var params getEditTodo

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
	todoID, _ := uuid.Parse(params.ID)
	todo := todos.Get(todoID)

	return handlers.CaseResponse(
		c,
		views.EditTodoView(*todo),
		nil,
		handlers.FullPageRender("EDIT TODO", views.EditTodoView(*todo)),
	)
}

type EditTodo struct {
	ID          string `uri:"id" validate:"required,uuid"`
	Description string `form:"description" validate:"required"`
	Completed   bool   `form:"completed" validate:"required"`
}

func (h *handler) EditTodo(c fiber.Ctx) error {
	var request EditTodo
	// time.Sleep(2 * time.Second)

	if err := c.Bind().URI(&request); err != nil {
		fmt.Println("test")
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}
	if err := c.Bind().Form(&request); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	stringJson, _ := json.Marshal(request)
	var result EditTodo
	json.Unmarshal(stringJson, &result)

	todoID, _ := uuid.Parse(result.ID)
	updatedTodo := todos.Update(todoID, result.Completed, result.Description)

	if handlers.IsHtmx(c) {
		// Check if update comes from edit page and set the Htmx location header for client redirect
		targetHeader := c.Get(util.HtmxTargetKey)
		if targetHeader == "app-content" {
			handlers.SetHtmxLocationHeader(c, "/todos", "#app-content")
		}
	}

	if !handlers.IsHtmx(c) {
		return c.Redirect().To("/todos")
	}

	return handlers.CaseResponse(
		c,
		components.Todo(*updatedTodo),
		updatedTodo,
		nil,
	)
}

type searchTodoRequest struct {
	Search string `form:"search"`
}

func (h *handler) SearchTodo(c fiber.Ctx) error {
	var req searchTodoRequest
	if err := c.Bind().Query(&req); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}
	result := utils.CopyString(req.Search)
	fmt.Println(req.Search)
	var foundTodos db.Todos = todos.Search(result)

	return handlers.CaseResponse(
		c,
		components.Todos(foundTodos),
		foundTodos,
		h.fullPageTodoComponent(foundTodos),
	)

}
