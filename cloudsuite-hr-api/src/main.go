package main

import (
	"log"

	"cloudsuite-hr-api/controllers"
	"cloudsuite-hr-api/database"
	"cloudsuite-hr-api/repositories"
	"cloudsuite-hr-api/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Create the repository, service, and controller
	timeRepository := repositories.NewTimeRepository(database.DB)
	timeService := services.NewTimeService(timeRepository)
	timeController := controllers.NewTimeController(timeService)

	// Create a new Fiber instance
	app := fiber.New()

	// Use Fiber's built-in logger middleware
	app.Use(logger.New())

	// Set up routes for the controller
	timeController.SetupRoutes(app)

	// Start the Fiber app and handle any errors during startup
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
