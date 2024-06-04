package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RequestLogger(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	stop := time.Now()
	latency := stop.Sub(start)

	method := c.Method()
	path := c.Path()
	status := c.Response().StatusCode()
	clientIP := c.IP()

	log.Printf("%s %s %d %s %s", method, path, status, latency, clientIP)

	return err
}
