package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// LoggingMiddleware configures structured logging for HTTP requests
func LoggingMiddleware() echo.MiddlewareFunc {
	return middleware.Logger()
}
