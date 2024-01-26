package api

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/luist1228/go-htmx-examples/db"
	"github.com/luist1228/go-htmx-examples/views"
	"github.com/luist1228/go-htmx-examples/views/layouts"
)

type Server struct {
	app *fiber.App
}

func NewServer() (*Server, error) {
	server := &Server{}
	server.setupRouter()
	return server, nil
}

func (s *Server) setupRouter() {
	app := fiber.New()

	app.Static("/assets", "./assets")

	db.FillTodos()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))

	api := app.Group("/api")

	app.Get("/", func(c fiber.Ctx) error {
		return Render(c, layouts.Main("Home", views.Home()))
	})

	api.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello from Api")
	})

	app.Use(NotFoundMiddleware)

	s.app = app
}

func (s *Server) Start(addr string) error {
	return s.app.Listen(addr)
}

func NotFoundMiddleware(c fiber.Ctx) error {
	return Render(c, layouts.Main("Not Found", views.NotFound()), templ.WithStatus(http.StatusNotFound))
}
