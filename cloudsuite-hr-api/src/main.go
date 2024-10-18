package main

import (
	"cloudsuite-hr-api/controllers"
	"cloudsuite-hr-api/database"
	_ "cloudsuite-hr-api/docs" // Importa a documentação gerada pelo swag
	"cloudsuite-hr-api/migrations"
	"cloudsuite-hr-api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title Cloudsuite HR API
// @version 1.0
// @description API para controle de registro de ponto.
// @host localhost:3000
// @BasePath /

func main() {
	db := database.InitDB()
	migrations.Migrate(db)

	timeService := services.NewTimeService(db)
	timeController := controllers.NewTimeController(timeService)

	app := fiber.New()
	app.Use(logger.New())

	// Rota do Swagger UI
	app.Get("/swagger/*", swagger.HandlerDefault)

	timeController.SetupRoutes(app)

	app.Listen(":3000")
}
