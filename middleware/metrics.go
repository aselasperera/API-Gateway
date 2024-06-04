package middleware

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
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
		req := convertRequest(c.Request())
		rw := &fiberToHTTPResponseWriter{c}
		promhttp.Handler().ServeHTTP(rw, req)
		return nil
	}
}

type fiberToHTTPResponseWriter struct {
	*fiber.Ctx
}

func (fw *fiberToHTTPResponseWriter) Header() http.Header {
	header := http.Header{}
	fw.Ctx.Response().Header.VisitAll(func(key, value []byte) {
		header.Set(string(key), string(value))
	})
	return header
}

func (fw *fiberToHTTPResponseWriter) Write(b []byte) (int, error) {
	return fw.Ctx.Write(b)
}

func (fw *fiberToHTTPResponseWriter) WriteHeader(statusCode int) {
	fw.Ctx.Status(statusCode)
}

func convertRequest(req *fasthttp.Request) *http.Request {
	uri := req.URI()
	url := &url.URL{
		Scheme:   string(uri.Scheme()),
		Host:     string(uri.Host()),
		Path:     string(uri.Path()),
		RawQuery: string(uri.QueryString()),
	}

	httpReq := &http.Request{
		Method: string(req.Header.Method()),
		URL:    url,
		Header: make(http.Header),
		Body:   io.NopCloser(bytes.NewReader(req.Body())),
	}

	req.Header.VisitAll(func(key, value []byte) {
		httpReq.Header.Set(string(key), string(value))
	})

	return httpReq
}
