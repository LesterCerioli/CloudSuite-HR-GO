package routes

import (
	"cloudsuite-hr-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, timeController *controllers.TimeController) {
	app.Post("/times", timeController.CreateTime)
	app.Get("/times", timeController.GetAllTimes)
}
