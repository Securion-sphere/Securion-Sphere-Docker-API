package handler

import (
	"errors"
	"net/http"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/usecase"
	imageErrors "github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/usecase/errors"
	"github.com/labstack/echo/v4"
)

type ImageHandler struct {
	imageUseCase *usecase.ImageUseCase
}

func NewImageHandler(uc *usecase.ImageUseCase) *ImageHandler {
	return &ImageHandler{imageUseCase: uc}
}

func (h *ImageHandler) UploadImage(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer src.Close()

	image, err := h.imageUseCase.UploadImage(c.Request().Context(), src)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, image)
}

func (h *ImageHandler) GetAllImage(c echo.Context) error {
	images, err := h.imageUseCase.GetAllImage(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, images)
}

func (h *ImageHandler) GetImage(c echo.Context) error {
	image, err := h.imageUseCase.GetImage(c.Request().Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, imageErrors.ErrImageNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, image)
}

func (h *ImageHandler) DeleteImage(c echo.Context) error {
	untags, err := h.imageUseCase.DeleteImage(c.Request().Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, imageErrors.ErrImageNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, untags)
}
