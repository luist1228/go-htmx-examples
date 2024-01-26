package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
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

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))

	api := app.Group("/api")

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	api.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello from Api")
	})

	s.app = app
}

func (s *Server) Start(addr string) error {
	return s.app.Listen(addr)
}
