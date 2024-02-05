package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/luist1228/go-htmx-examples/templates/views"
	"github.com/luist1228/go-htmx-examples/templates/views/layouts"
)

func (h *Handler) Register(app *fiber.App) {
	api := app.Group("/api")

	app.Get("/", func(c fiber.Ctx) error {
		return Render(c, FullPageRender("Home", views.Home()))
	})

	app.Get("/test", func(c fiber.Ctx) error {
		return Render(c, views.Home())
	})

	api.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello from Api")
	})

	app.Use(NotFoundMiddleware)
}

func NotFoundMiddleware(c fiber.Ctx) error {
	return Render(c, layouts.Main("Not Found", views.NotFound()), templ.WithStatus(http.StatusNotFound))
}

func FullPageRender(title string, content templ.Component) templ.Component {
	return layouts.Main(title, layouts.App(content))
}
