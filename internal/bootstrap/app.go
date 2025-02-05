package bootstrap

import (
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/adapters/in/http/handler"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/adapters/in/http/routes"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/adapters/out/docker"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/usecase"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/infrastructure"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/middleware"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Bootstrap initializes and returns the Echo server
func Bootstrap() (*echo.Echo, *config.Config) {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// logger.SetLevel(log.INFO)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error("Failed to load config: ", err)
	}

	//After router
	e.Use(middleware.EchoCorsMiddleware())
	e.Use(middleware.LoggingMiddleware())

	// Create a shared Docker client
	dockerClient, err := infrastructure.NewDockerClient()
	if err != nil {
		log.Fatal("Failed to initialize Docker client:", err)
		return nil, cfg
	}

	// Inject dependencies
	containerAdapter := docker.NewDockerContainerAdapter(dockerClient)
	containerUseCase := usecase.NewContainerUseCase(containerAdapter)
	containerHandler := handler.NewContainerHandler(containerUseCase)

	infoAdapter := docker.NewDockerInfoAdapter(dockerClient)
	infoUseCase := usecase.NewInfoUsecase(infoAdapter)
	healthHandler := handler.NewHealthHandler(infoUseCase)

	// Register Routes
	routes.RegisterContainerRoutes(e, containerHandler)
	routes.RegisterHealthRoute(e, healthHandler)

	// Log server info
	log.Info("Server initialized on port:", cfg.AppPort)

	return e, cfg
}
