package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func MetricsMiddleware(c *fiber.Ctx) error {
	httpRequestsTotal.WithLabelValues(c.Path()).Inc()
	return c.Next()
}

func PrometheusHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		promhttp.Handler().ServeHTTP(c.Response().BodyWriter(), c.Request())
		return nil
	}
}
