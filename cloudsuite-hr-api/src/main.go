package main

import (
	"cloudsuite-hr-api/config"
	"cloudsuite-hr-api/controllers"
	"cloudsuite-hr-api/repositories"
	"cloudsuite-hr-api/routes"
	"cloudsuite-hr-api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
	config.InitDB()

	timeRespository := repositories.NewTimeRepository(config.DB)
	timeService := services.NewTimeService(timeRespository)
	timeController := controllers.NewTimeController(timeService)

	app := fiber.New()

	app.Use(logger.New())

	routes.SetupRoutes(app, timeController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
