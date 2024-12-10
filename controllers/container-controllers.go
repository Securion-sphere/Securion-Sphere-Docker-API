package controllers

import (
	"docker-api/domain"
	"docker-api/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ContainerController struct {
	containerService domain.ContainerService
}

func NewContainerController(service domain.ContainerService) *ContainerController {
	return &ContainerController{containerService: service}
}

func (cc *ContainerController) GetContainer(ctx echo.Context) error {
	containers, err := cc.containerService.GetContainer()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}
	return ctx.JSON(http.StatusOK, containers)
}

func (cc *ContainerController) CreateContainer(ctx echo.Context) error {
	container := dto.CreateContainerDto{}
	validator := validator.New()

	err := ctx.Bind(&container)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"msg": err.Error()})
	}
	err = validator.Struct(&container)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"msg": err.Error()})
	}

	createdContainer, err := cc.containerService.CreateContainer(container)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, *createdContainer)
}

func (cc *ContainerController) DeleteContainer(c echo.Context) error {
	container := dto.DeleteContainerDto{}
	validator := validator.New()

	err := c.Bind(&container)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": err.Error()})
	}
	err = validator.Struct(&container)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"msg": err.Error()})
	}

	res, err := cc.containerService.DeleteContainer(container)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
