package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/crontastic/app/models"
	"github.com/mrrizkin/crontastic/system/types"
)

func (h *Handlers) UserPage(c *fiber.Ctx) error {
	return c.Render("pages/user", fiber.Map{})
}

func (h *Handlers) UserCreate(c *fiber.Ctx) error {
	var (
		err     error
		payload models.User
	)

	err = c.BodyParser(&payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	user, err := h.userService.Create(&payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed create user",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success create user",
		Data:    user,
	})
}

func (h *Handlers) UserFindAll(c *fiber.Ctx) error {
	var (
		err   error
		users []models.User
	)

	users, err = h.userService.FindAll()
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed get users",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get users",
		Data:    users,
	})
}

func (h *Handlers) UserFindByID(c *fiber.Ctx) error {
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

	user, err := h.userService.FindByID(uint(id))
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed get user",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get user",
		Data:    user,
	})
}

func (h *Handlers) UserUpdate(c *fiber.Ctx) error {
	var (
		err     error
		id      int
		payload models.User
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

	user, err := h.userService.Update(uint(id), &payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed update user",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success update user",
		Data:    user,
	})
}

func (h *Handlers) UserDelete(c *fiber.Ctx) error {
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

	err = h.userService.Delete(uint(id))
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed delete user",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success delete user",
	})
}
