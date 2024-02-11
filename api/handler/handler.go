package handler

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/luist1228/go-htmx-examples/db"
)

const (
	htmxHeaderKey = "HX-Request"
)

var validate = validator.New()

type Handler struct {
	todos    db.Todos
	validate *validator.Validate
}

func NewHandler() *Handler {
	todos := db.FillTodos()
	return &Handler{
		todos:    *todos,
		validate: validate,
	}
}

func isHtmx(c fiber.Ctx) bool {
	return c.Get(htmxHeaderKey) == "true"
}

func isApiRequest(c fiber.Ctx) bool {
	return strings.Contains(c.Path(), "/api")
}
