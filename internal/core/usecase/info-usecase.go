package usecase

import (
	"context"
	"net/http"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/ports"
	"github.com/labstack/echo/v4"
)

type InfoUseCase struct {
	infoService ports.InfoService
}

func NewInfoUsecase(infoService ports.InfoService) *InfoUseCase {
	return &InfoUseCase{infoService: infoService}
}

func (uc *InfoUseCase) Info(ctx context.Context) error {
	if _, err := uc.infoService.GetInfo(ctx); err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, err)
	}
	return nil
}
