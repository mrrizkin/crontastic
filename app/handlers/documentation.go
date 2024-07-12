package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/crontastic/resources/views/pages"
)

func (h *Handlers) DocumentationPage(c *fiber.Ctx) error {
  return h.Render(c, pages.Docs())
}
