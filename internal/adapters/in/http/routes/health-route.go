package routes

import (
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/adapters/in/http/handler"
	"github.com/labstack/echo/v4"
)

func RegisterHealthRoute(e *echo.Echo, h *handler.HealthHandler) {
	e.GET("/health", h.HealthCheck)
}
