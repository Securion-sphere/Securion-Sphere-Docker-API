package routes

import (
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/adapters/in/http/handler"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterContainerRoute(e *echo.Echo, h *handler.ContainerHandler) {
	g := e.Group("/container", middleware.EchoJWTMiddleware())

	g.POST("", h.CreateContainer)
	g.GET("", h.ListContainers)
	g.GET("/:id", h.GetContainer)
	g.DELETE("/:id", h.StopContainer)
}
