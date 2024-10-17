package main

import (
	"cloudsuite-hr-api/controllers"
	"cloudsuite-hr-api/database"
	"cloudsuite-hr-api/migrations" // Adicione esta importação
	"cloudsuite-hr-api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db := database.InitDB() // Inicializa o banco de dados
	migrations.Migrate(db)  // Chama a migração

	timeService := services.NewTimeService(db) // Passa a instância do db
	timeController := controllers.NewTimeController(timeService)

	app := fiber.New()
	app.Use(logger.New())

	timeController.SetupRoutes(app)

	app.Listen(":3000")
}
