package ports

import (
	"context"
	"mime/multipart"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/domain"
	"github.com/docker/docker/api/types/image"
)

type ImageService interface {
	LoadImage(ctx context.Context, file multipart.File) (*image.LoadResponse, error)
	ListImage(ctx context.Context, listOptions *image.ListOptions) ([]domain.Image, error)
	RemoveImage(
		ctx context.Context,
		id string,
		removeOptions *image.RemoveOptions,
	) ([]image.DeleteResponse, error)
	InspectImage(
		ctx context.Context,
		id string,
	) (*domain.Image, error)
}
