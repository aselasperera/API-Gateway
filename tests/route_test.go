package tests

import (
	"api-gateway/routes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserRoute(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	req := httptest.NewRequest("GET", "/api/v1/users", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestProductRoute(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	req := httptest.NewRequest("GET", "/api/v1/products", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
