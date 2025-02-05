package routes

import (
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/adapters/in/http/handler"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterImageRoute(e *echo.Echo, h *handler.ImageHandler) {
	g := e.Group("/image")
	g.Use(middleware.EchoJWTMiddleware())

	g.POST("", h.UploadImage)
	g.GET("", h.GetAllImage)
	g.GET("/:id", h.GetImage)
	g.DELETE("/:id", h.DeleteImage)
}
