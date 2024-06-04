package routes

import (
	"api-gateway/config"
	"api-gateway/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func SetupRoutes(app *fiber.App) {
	app.Use(middleware.AuthMiddleware)
	app.Use(middleware.RateLimiter())
	app.Use(middleware.RequestLogger)
	app.Use(middleware.MetricsMiddleware)
	app.Use(middleware.ValidateRequestBody)

	app.All("/api/v1/users/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, config.UserServiceURL)
	})

	app.All("/api/v1/products/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, config.ProductServiceURL)
	})

	app.Get("/metrics", middleware.PrometheusHandler())
}
