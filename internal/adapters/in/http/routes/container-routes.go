package routes

import (
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/adapters/in/http/handler"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterContainerRoutes(e *echo.Echo, h *handler.ContainerHandler) {
	r := e.Group("/container")

	r.Use(middleware.EchoJWTMiddleware())

	r.POST("", h.CreateContainer)
	r.GET("", h.ListContainers)
	r.GET("/:id", h.GetContainer)
	r.DELETE("/:id", h.StopContainer)
}
