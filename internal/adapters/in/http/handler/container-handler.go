package handler

import (
	"net/http"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/adapters/in/http/handler/dto"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/usecase"
	"github.com/labstack/echo/v4"
)

type ContainerHandler struct {
	containerUseCase *usecase.ContainerUseCase
}

// NewContainerHandler initializes a new handler
func NewContainerHandler(uc *usecase.ContainerUseCase) *ContainerHandler {
	return &ContainerHandler{containerUseCase: uc}
}

// CreateContainer handles the creation of a container
func (h *ContainerHandler) CreateContainer(c echo.Context) error {
	request := dto.CreateContainerDto{}
	if err := c.Bind(&request); err != nil {
		return err
	}
	if err := c.Validate(&request); err != nil {
		return err
	}

	container, err := h.containerUseCase.CreateContainer(c.Request().Context(), request.Image, request.ContainerPort, request.HostPort, request.Flag)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, container)
}

func (h *ContainerHandler) ListContainers(c echo.Context) error {
	containers, err := h.containerUseCase.GetAllContainer(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, containers)
}

func (h *ContainerHandler) GetContainer(c echo.Context) error {
	conatiner, err := h.containerUseCase.GetContainer(c.Request().Context(), c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, conatiner)
}

func (h* ContainerHandler) StopContainer(c echo.Context) error {
	err := h.containerUseCase.StopContainer(c.Request().Context(), c.Param("id"))
	if err != nil {

	}
	return c.JSON(http.StatusOK, map[string]interface{}{"msg": `${c.Param("id")}`})
}