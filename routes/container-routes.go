package routes

import (
	"docker-api/controllers"
	"docker-api/service"

	"github.com/labstack/echo/v4"
)

func ContainerRoutes(g *echo.Group) {
	containerService := service.NewContainerService()
	containerController := controllers.NewContainerController(containerService)

	g.GET("/container", containerController.GetContainer)
	g.POST("/container", containerController.CreateContainer)
	g.DELETE("/container", containerController.DeleteContainer)
}
