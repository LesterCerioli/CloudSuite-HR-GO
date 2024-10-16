package main

import (
	"cloudsuite-hr-api/controllers"
	"cloudsuite-hr-api/database"
	"cloudsuite-hr-api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.InitDB()

	timeService := services.NewTimeService()
	timeController := controllers.NewTimeController(timeService, database.InitDB())

	app := fiber.New()

	app.Use(logger.New())

	timeController.SetupRoutes(app)

	app.Listen(":3000")
}
