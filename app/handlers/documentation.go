package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handlers) DocumentationPage(c *fiber.Ctx) error {
	return c.Render("pages/docs", fiber.Map{})
}
