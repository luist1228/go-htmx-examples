package handlers

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/luist1228/go-htmx-examples/templates/views"
	"github.com/luist1228/go-htmx-examples/templates/views/layouts"
)

const (
	htmxHeaderKey = "HX-Request"
)

func NotFoundMiddleware(c fiber.Ctx) error {
	return Render(c, layouts.Main("Not Found", views.NotFound()), templ.WithStatus(http.StatusNotFound))
}

func FullPageRender(title string, content templ.Component) templ.Component {
	return layouts.Main(title, layouts.App(content))
}

func IsHtmx(c fiber.Ctx) bool {
	return c.Get(htmxHeaderKey) == "true"
}

func IsApiRequest(c fiber.Ctx) bool {
	return strings.Contains(c.Path(), "/api")
}

func CaseResponse(
	c fiber.Ctx,
	htmxContent templ.Component,
	apiData any,
	fullPage templ.Component,
) error {
	if IsApiRequest(c) {
		return c.JSON(apiData)
	}

	if IsHtmx(c) {
		return Render(c, htmxContent)
	}

	return Render(c, fullPage)
}
