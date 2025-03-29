package usecase

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/domain"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/ports"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/usecase/errors"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
)

type ImageUseCase struct {
	imageService ports.ImageService
}

func NewImageUseCase(imageService ports.ImageService) *ImageUseCase {
	return &ImageUseCase{imageService: imageService}
}

func (uc *ImageUseCase) GetAllImage(ctx context.Context) ([]domain.Image, error) {
	images, err := uc.imageService.ListImage(ctx, &image.ListOptions{})
	if err != nil {
		return nil, err
	}
	return images, nil
}

func (uc *ImageUseCase) GetImage(ctx context.Context, id string) (*domain.Image, error) {
	image, err := uc.imageService.InspectImage(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "No such image") {
			return nil, errors.ErrImageNotFound
		}
		return nil, err
	}
	return image, nil
}

func (uc *ImageUseCase) UploadImage(
	ctx context.Context,
	file multipart.File,
) (*domain.Image, error) {
	resp, err := uc.imageService.LoadImage(ctx, file)
	if err != nil {
		return nil, err
	}
	// defer resp.Body.Close()
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			err = fmt.Errorf("failed to close file: %w", closeErr) // Capture the error
		}
	}()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respString := string(respBytes)

	if !strings.Contains(respString, "Loaded image") {
		return nil, errors.ErrImageLoadFailed
	}

	parts := strings.SplitN(respString, "Loaded image: ", 2)
	name := strings.SplitN(parts[1], ":", 2)[0]

	images, err := uc.imageService.ListImage(ctx, &image.ListOptions{Filters: filters.NewArgs(
		filters.KeyValuePair{Key: "reference", Value: name},
	)})
	if err != nil {
		return nil, errors.ErrFailedToListAfterLoad
	}

	if len(images) > 1 {
		return nil, errors.ErrMultipleImages
	} else if len(images) == 0 {
		return nil, errors.ErrImageLoadFailed
	}
	return &domain.Image{ID: images[0].ID, RepoTags: images[0].RepoTags, Size: images[0].Size}, nil
}

func (uc *ImageUseCase) DeleteImage(
	ctx context.Context,
	id string,
) (map[string]interface{}, error) {
	var resp []image.DeleteResponse

	resp, err := uc.imageService.RemoveImage(ctx, id, &image.RemoveOptions{})
	if err != nil {
		resp, err = uc.imageService.RemoveImage(ctx, id, &image.RemoveOptions{Force: true})
		if err != nil {
			return nil, err
		}
	}

	if len(resp) == 0 {
		return nil, errors.ErrImageNotFound
	}

	var removedId []string
	for _, r := range resp {
		removedId = append(removedId, r.Untagged)
	}

	return map[string]interface{}{"Removed Image ID": removedId}, nil
}
