package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handlers) ManualTasksPage(c *fiber.Ctx) error {
	return c.Render("pages/manual_tasks", fiber.Map{})
}
