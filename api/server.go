package api

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/luist1228/go-htmx-examples/api/handler"
)

const (
	htmxHeaderKey = "HX-Request"
)

type Server struct {
	app     *fiber.App
	handler handler.Handler
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewServer() (*Server, error) {
	handler := handler.NewHandler()
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

	server := &Server{
		handler: *handler,
		app:     app,
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

	s.handler.Register(s.app)
}

func (s *Server) Start(addr string) error {
	return s.app.Listen(addr)
}
