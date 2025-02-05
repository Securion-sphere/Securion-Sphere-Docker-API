package ports

import (
	"context"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/domain"
	"github.com/docker/docker/api/types/container"
)

type ContainerService interface {
	CreateContainer(
		ctx context.Context,
		config *container.Config,
		hostConfig *container.HostConfig,
	) (string, error)
	ListContainers(
		ctx context.Context,
		listOption *container.ListOptions,
	) ([]domain.Container, error)
	StartContainer(ctx context.Context, id string, startOptions *container.StartOptions) error
	StopContainer(ctx context.Context, id string, stopOptions *container.StopOptions) error
	RemoveContainer(ctx context.Context, id string, removeOptions *container.RemoveOptions) error
}
