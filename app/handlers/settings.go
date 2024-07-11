package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handlers) SettingsPage(c *fiber.Ctx) error {
	return c.Render("pages/settings", fiber.Map{})
}
