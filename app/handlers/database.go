package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/crontastic/app/models"
	"github.com/mrrizkin/crontastic/system/types"
)

func (h *Handlers) DatabaseCreate(c *fiber.Ctx) error {
	var (
		err     error
		payload models.Database
	)

	err = c.BodyParser(&payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	database, err := h.databaseService.Create(&payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed create database",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success create database",
		Data:    database,
	})
}

func (h *Handlers) DatabaseFindAll(c *fiber.Ctx) error {
	var (
		err       error
		databases []models.Database
	)

	databases, err = h.databaseService.FindAll()
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed get databases",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get databases",
		Data:    databases,
	})
}

func (h *Handlers) DatabaseFindByID(c *fiber.Ctx) error {
	var (
		err error
		id  int
	)

	id, err = c.ParamsInt("id")
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	database, err := h.databaseService.FindByID(uint(id))
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed get database",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get database",
		Data:    database,
	})
}

func (h *Handlers) DatabaseUpdate(c *fiber.Ctx) error {
	var (
		err     error
		id      int
		payload models.Database
	)

	id, err = c.ParamsInt("id")
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	err = c.BodyParser(&payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	database, err := h.databaseService.Update(uint(id), &payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed update database",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success update database",
		Data:    database,
	})
}

func (h *Handlers) DatabaseDelete(c *fiber.Ctx) error {
	var (
		err error
		id  int
	)

	id, err = c.ParamsInt("id")
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "id not valid",
			Debug:   err.Error(),
		}, 400)
	}

	err = h.databaseService.Delete(uint(id))
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed delete database",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success delete database",
	})
}
