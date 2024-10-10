package controllers

import (
	"cloudsuite-hr-api/models"
	"cloudsuite-hr-api/services"
	"github.com/gofiber/fiber/v2"
)

type TimeController struct {
	service services.TimeService
}

func NewTimeController(service services.TimeService) *TimeController {
	return &TimeController{service: service}
}

func (c *TimeController) CreateTime(ctx *fiber.Ctx) error {
	var time models.Time
	if err := ctx.BodyParser(&time); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if err := c.service.CreateTime(time); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create time entry",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(time)
}

func (c *TimeController) GetAllTimes(ctx *fiber.Ctx) error {
	times, err := c.service.GetAllTimes()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get times",
		})
	}
	return ctx.JSON(times)
}
