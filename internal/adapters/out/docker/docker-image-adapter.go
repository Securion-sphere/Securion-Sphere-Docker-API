package docker

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/domain"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/ports"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/infrastructure"
	"github.com/docker/docker/api/types/image"
)

type DockerImageAdapter struct {
	dockerClient *infrastructure.DockerClient
}

var _ ports.ImageService = (*DockerImageAdapter)(nil)

func NewDockerImageAdapter(cli *infrastructure.DockerClient) *DockerImageAdapter {
	return &DockerImageAdapter{dockerClient: cli}
}

func (a *DockerImageAdapter) LoadImage(
	ctx context.Context,
	file multipart.File,
) (*image.LoadResponse, error) {
	resp, err := a.dockerClient.Client.ImageLoad(ctx, file, false)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (a *DockerImageAdapter) ListImage(
	ctx context.Context,
	listOptions *image.ListOptions,
) ([]domain.Image, error) {
	images, err := a.dockerClient.Client.ImageList(ctx, *listOptions)
	if err != nil {
		return nil, err
	}

	var result []domain.Image
	for _, i := range images {
		var id string

		if _, err := fmt.Sscanf(i.ID, "sha256:%s", &id); err != nil {
			return nil, err
		}
		result = append(result, domain.Image{ID: id, RepoTags: i.RepoTags, Size: i.Size})
	}
	return result, nil
}

func (a *DockerImageAdapter) InspectImage(
	ctx context.Context,
	id string,
) (*domain.Image, error) {
	image, _, err := a.dockerClient.Client.ImageInspectWithRaw(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.Image{ID: image.ID, RepoTags: image.RepoTags, Size: image.Size}, nil
}

func (a *DockerImageAdapter) RemoveImage(
	ctx context.Context,
	id string,
	removeOptions *image.RemoveOptions,
) ([]image.DeleteResponse, error) {
	resp, err := a.dockerClient.Client.ImageRemove(ctx, id, *removeOptions)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
