package api

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/luist1228/go-htmx-examples/api/handlers"
	"github.com/luist1228/go-htmx-examples/api/handlers/todos"
)

const (
	htmxHeaderKey = "HX-Request"
)

type Server struct {
	app         *fiber.App
	todoHandler todos.Handler
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var validate = validator.New()

func NewServer() (*Server, error) {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := http.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			c.Set("Content-Type", "text/plain")
			return c.Status(code).SendString(err.Error())
		},
	})

	todosHandler := todos.NewHandler(validate)
	server := &Server{
		app:         app,
		todoHandler: todosHandler,
	}

	server.setupRouter()
	return server, nil
}

func (s *Server) setupRouter() {

	s.app.Static("/assets", "./assets")

	s.app.Use(logger.New())

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))

	// Check if request has HTMX headers and add it to the context
	s.app.Use(func(c fiber.Ctx) error {
		htmxHeader := c.Get(htmxHeaderKey)
		c.Context().SetUserValue("isHtmx", htmxHeader == "true")
		return c.Next()
	})

	api := s.app.Group("/api")

	s.app.Get("/", func(c fiber.Ctx) error {
		return c.Redirect().To("/todos")
	})

	// Setup Todos Routes
	s.todoHandler.Mount(s.app, api)

	api.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello from Api")
	})

	s.app.Use(handlers.NotFoundMiddleware)
}

func (s *Server) Start(addr string) error {
	return s.app.Listen(addr)
}
