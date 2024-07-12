package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/crontastic/resources/views/pages"
)

func (h *Handlers) SettingsPage(c *fiber.Ctx) error {
	return h.Render(c, pages.Settings())
}
