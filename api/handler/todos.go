package handler

import "github.com/gofiber/fiber/v3"

func (h *Handler) Todos(c fiber.Ctx)error {

	return c.JSON(h.todos)
	
}
