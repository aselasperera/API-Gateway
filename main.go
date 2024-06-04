package main

import (
	"api-gateway/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Register middleware
	app.Use(middleware.MetricsMiddleware)

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the API Gateway!")
	})

	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Users Service"})
	})
	app.Get("/api/v1/products", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Products Service"})
	})

	// Prometheus metrics endpoint
	app.Get("/metrics", middleware.PrometheusHandler())

	// Start server
	log.Fatal(app.Listen(":3001"))
}
