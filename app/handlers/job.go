package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/crontastic/app/models"
	"github.com/mrrizkin/crontastic/system/types"
)

func (h *Handlers) JobCreate(c *fiber.Ctx) error {
	var (
		err     error
		payload models.Job
	)

	err = c.BodyParser(&payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "payload not valid",
			Debug:   err.Error(),
		}, 400)
	}

	job, err := h.jobService.Create(&payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed create job",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success create job",
		Data:    job,
	})
}

func (h *Handlers) JobFindAll(c *fiber.Ctx) error {
	var (
		err  error
		jobs []models.Job
	)

	jobs, err = h.jobService.FindAll()
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed get jobs",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get jobs",
		Data:    jobs,
	})
}

func (h *Handlers) JobFindByID(c *fiber.Ctx) error {
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

	job, err := h.jobService.FindByID(uint(id))
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed get job",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success get job",
		Data:    job,
	})
}

func (h *Handlers) JobUpdate(c *fiber.Ctx) error {
	var (
		err     error
		id      int
		payload models.Job
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

	job, err := h.jobService.Update(uint(id), &payload)
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed update job",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success update job",
		Data:    job,
	})
}

func (h *Handlers) JobDelete(c *fiber.Ctx) error {
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

	err = h.jobService.Delete(uint(id))
	if err != nil {
		return h.SendJson(c, types.Response{
			Success: false,
			Message: "failed delete job",
			Debug:   err.Error(),
		}, 500)
	}

	return h.SendJson(c, types.Response{
		Success: true,
		Message: "success delete job",
	})
}
