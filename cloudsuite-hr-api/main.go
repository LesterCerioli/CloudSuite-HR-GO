package main

import (
	"cloudsuite-hr-api/config"
	"cloudsuite-hr-api/controllers"
	"cloudsuite-hr-api/repositories"
	"cloudsuite-hr-api/routes"
	"cloudsuite-hr-api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.InitDB()

	timeRespository := repositories.NewTimeRepository(config.DB)
	timeService := services.NewTimeService(timeRespository)
	timeController := controllers.NewTimeController(timeService)

	app := fiber.New()

	app.Use(logger.New())

	routes.SetupRoutes(app, timeController)

	app.Listen(":8080")
}
