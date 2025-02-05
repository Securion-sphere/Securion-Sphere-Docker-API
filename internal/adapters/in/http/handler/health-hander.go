package handler

import (
	"net/http"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/usecase"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	infoUseCase *usecase.InfoUseCase
}

func NewHealthHandler(infoUseCase *usecase.InfoUseCase) *HealthHandler {
	return &HealthHandler{infoUseCase: infoUseCase}
}

func (h *HealthHandler) HealthCheck(c echo.Context) error {
	type HealthCheckResponse struct {
		Status string
	}

	if err := h.infoUseCase.Info(c.Request().Context()); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &HealthCheckResponse{Status: "ok"})
}
