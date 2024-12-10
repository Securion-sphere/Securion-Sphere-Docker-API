package main

import (
	"docker-api/configs"
	"docker-api/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	}))

	// Load configuration
	config, err := configs.LoadConfig(".")
	if err != nil {
		e.Logger.Fatal("cannot load config:", err)
	}

	api := e.Group("/" + config.DockerAPIGroup)

	// Routes
	routes.ContainerRoutes(api)
	routes.ImageRoutes(api)

	// Start server
	e.Logger.Fatal(e.Start(":" + string(config.AppPort)))
}
