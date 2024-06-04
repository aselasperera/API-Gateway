package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Users Service"})
	})

	log.Fatal(app.Listen(":3001"))
}
