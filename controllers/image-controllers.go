package controllers

import (
	"docker-api/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ImageController struct {
	imageService domain.ImageService
}

func NewImageController(service domain.ImageService) *ImageController {
	return &ImageController{imageService: service}
}

func (c *ImageController) GetImage(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"msg": "Get Image"})
}

func (c *ImageController) LoadImage(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"msg": "Load Image"})
}

func (c *ImageController) DeleteImage(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"msg": "Delete Image"})
}
