package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) IndexPage(c *fiber.Ctx) error {
	return c.Render("pages/welcome", fiber.Map{
		"csrf": c.Locals("X-CSRF-Token"),
	})
}
