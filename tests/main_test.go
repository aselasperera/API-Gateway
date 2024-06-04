package tests

import (
	"api-gateway/config"
	"api-gateway/middleware"
	"api-gateway/routes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestMainApp(t *testing.T) {
	config.LoadConfig()

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	routes.SetupRoutes(app)

	req := httptest.NewRequest("GET", "/api/v1/users", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	req = httptest.NewRequest("GET", "/api/v1/products", nil)
	resp, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
