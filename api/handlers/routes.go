package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/luist1228/go-htmx-examples/templates/views"
	"github.com/luist1228/go-htmx-examples/templates/views/layouts"
	"github.com/luist1228/go-htmx-examples/util"
)


func NotFoundMiddleware(c fiber.Ctx) error {
	return Render(c, layouts.Main("Not Found", views.NotFound()), templ.WithStatus(http.StatusNotFound))
}

func FullPageRender(title string, content templ.Component) templ.Component {
	return layouts.Main(title, layouts.App(content))
}

func IsHtmx(c fiber.Ctx) bool {
	return c.Get(util.HtmxHeaderKey) == "true"
}

func IsApiRequest(c fiber.Ctx) bool {
	return strings.Contains(c.Path(), "/api")
}

func CaseResponse(
	c fiber.Ctx,
	partial templ.Component,
	apiData any,
	fullPage templ.Component,
) error {
	if IsApiRequest(c) {
		return c.JSON(apiData)
	}

	if IsHtmx(c) {
		return Render(c, partial)
	}

	return Render(c, fullPage)
}

func SetHtmxLocationHeader(c fiber.Ctx, path string, target string) error {
	location := struct {
		Path   string `json:"path"`
		Target string `json:"target"`
	}{
		Path:   path,
		Target: target,
	}

	l, err := json.Marshal(location)

	if err != nil {
		return err
	}
	c.Response().Header.Add(util.HtmxLocationKey, string(l))
	return nil
}
