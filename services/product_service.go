package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/v1/products", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Products Service"})
	})

	log.Fatal(app.Listen(":3002"))
}
