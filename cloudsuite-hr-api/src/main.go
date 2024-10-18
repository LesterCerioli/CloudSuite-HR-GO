package main

import (
	"cloudsuite-hr-api/controllers"
	"cloudsuite-hr-api/database"
	"cloudsuite-hr-api/migrations"
	"cloudsuite-hr-api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db := database.InitDB()
	migrations.Migrate(db)

	timeService := services.NewTimeService(db)
	timeController := controllers.NewTimeController(timeService)

	app := fiber.New()
	app.Use(logger.New())

	timeController.SetupRoutes(app)

	app.Listen(":3000")
}
