package api

import (
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

func NewServer() (*Server, error) {
	handler := handler.NewHandler()
	app := fiber.New()

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
