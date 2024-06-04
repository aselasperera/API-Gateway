package main

import (
	"api-gateway/config"
	"api-gateway/middleware"
	"api-gateway/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadConfig()

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Use(logger.New())

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
