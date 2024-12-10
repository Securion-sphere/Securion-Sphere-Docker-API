package routes

import (
	"docker-api/controllers"
	"docker-api/service"

	"github.com/labstack/echo/v4"
)

func ImageRoutes(g *echo.Group) {
	imageService := service.NewImageService()
	imageController := controllers.NewImageController(imageService)

	g.GET("/image", imageController.GetImage)
	g.POST("/image", imageController.LoadImage)
	g.DELETE("/image", imageController.DeleteImage)
}
