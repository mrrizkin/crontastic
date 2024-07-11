package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handlers) AutoTasksPage(c *fiber.Ctx) error {
	return c.Render("pages/auto_tasks", fiber.Map{})
}
