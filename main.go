package main

import (
	"docker-api/configs"
	"docker-api/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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
	e.Logger.SetLevel(log.INFO)

	e.Logger.Info("Starting the application...")

	// Load configuration
	config, err := configs.LoadConfig(e.Logger, ".")
	if err != nil {
		e.Logger.Fatal("Failed to load config:", err)
	}

	api := e.Group("/" + config.DockerAPIGroup)

	// Routes
	routes.ContainerRoutes(api)
	routes.ImageRoutes(api)

	// Start server
	e.Logger.Fatal(e.Start(":" + string(config.AppPort)))
}
