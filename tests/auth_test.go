package tests

import (
	"api-gateway/middleware"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	app := fiber.New()

	app.Use(middleware.AuthMiddleware)

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("success")
	})

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer mysecrettoken")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	req = httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer wrongtoken")
	resp, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)
}
