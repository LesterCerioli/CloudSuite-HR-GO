package controllers

import (
	"cloudsuite-hr-api/models"
	"cloudsuite-hr-api/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type TimeController struct {
	service services.TimeService
}

func NewTimeController(service services.TimeService) *TimeController {
	return &TimeController{
		service: service,
	}
}

func (c *TimeController) SetupRoutes(app *fiber.App) {
	app.Post("/times", c.CreateTime)
	app.Get("/times", c.GetAllTimes)
	app.Get("/times/date/:date", c.GetTimesByDate)
	app.Get("/times/year/:year", c.GetTimesByYear)
	app.Get("/times/month/:month", c.GetTimesByMonth)
	app.Get("/times/day/:day", c.GetTimesByDay)
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

func (c *TimeController) GetTimesByDate(ctx *fiber.Ctx) error {
	date := ctx.Params("date")
	times, err := c.service.GetTimesByDate(date)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get times by date",
		})
	}
	return ctx.JSON(times)
}

func (c *TimeController) GetTimesByYear(ctx *fiber.Ctx) error {
	year := ctx.Params("year")
	parsedYear, err := strconv.Atoi(year)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid year",
		})
	}
	times, err := c.service.GetTimesByYear(parsedYear)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get times by year",
		})
	}
	return ctx.JSON(times)
}

func (c *TimeController) GetTimesByMonth(ctx *fiber.Ctx) error {
	month := ctx.Params("month")
	parsedMonth, err := strconv.Atoi(month)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid month",
		})
	}
	times, err := c.service.GetTimesByMonth(parsedMonth)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get times by month",
		})
	}
	return ctx.JSON(times)
}

func (c *TimeController) GetTimesByDay(ctx *fiber.Ctx) error {
	day := ctx.Params("day")
	parsedDay, err := strconv.Atoi(day)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid day",
		})
	}
	times, err := c.service.GetTimesByDay(parsedDay)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get times by day",
		})
	}
	return ctx.JSON(times)
}
