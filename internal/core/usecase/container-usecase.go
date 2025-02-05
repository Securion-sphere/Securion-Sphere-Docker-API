package usecase

import (
	"context"
	"fmt"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/domain"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/ports"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/usecase/errors"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/go-connections/nat"
)

type ContainerUseCase struct {
	containerService ports.ContainerService
}

func NewContainerUseCase(containerService ports.ContainerService) *ContainerUseCase {
	return &ContainerUseCase{containerService: containerService}
}

func (uc *ContainerUseCase) CreateContainer(
	ctx context.Context,
	image string,
	containerPort uint16,
	hostPort uint16,
	flag string,
) (*domain.Container, error) {
	portBindings := nat.PortMap{
		nat.Port(fmt.Sprintf("%d", containerPort)): []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: fmt.Sprintf("%d", hostPort),
			},
		},
	}

	id, err := uc.containerService.CreateContainer(ctx, &container.Config{
		Image: image,
		Env:   []string{"SC_FLAG=" + flag},
		ExposedPorts: nat.PortSet{
			nat.Port(fmt.Sprintf("%d", containerPort)): struct{}{},
		},
	}, &container.HostConfig{
		PortBindings: portBindings,
		AutoRemove:   true,
	})
	if err != nil {
		return nil, errors.ErrFailedToList
	}

	err = uc.containerService.StartContainer(ctx, id, &container.StartOptions{})
	if err != nil {
		return nil, errors.ErrContainerStart
	}

	containers, err := uc.containerService.ListContainers(
		ctx,
		&container.ListOptions{
			All: true,
			Filters: filters.NewArgs(
				filters.KeyValuePair{Key: "id", Value: id},
			),
		},
	)
	if err != nil {
		return nil, errors.ErrFailedToListAfterStart
	}

	if len(containers) > 1 {
		for _, c := range containers {
			if err := uc.containerService.StopContainer(ctx, c.ID, &container.StopOptions{}); err != nil {
				return nil, errors.ErrContainerStopAfterStart
			}
		}
		return nil, errors.ErrMultipleContainers
	} else if len(containers) == 0 {
		return nil, errors.ErrContainerExited
	}

	return &domain.Container{
		ID:      containers[0].ID,
		Name:    containers[0].Name,
		Image:   containers[0].Image,
		Ports:   containers[0].Ports,
		Status:  containers[0].Status,
		Created: containers[0].Created,
	}, nil
}

func (uc *ContainerUseCase) GetAllContainer(ctx context.Context) ([]domain.Container, error) {
	containers, err := uc.containerService.ListContainers(ctx, &container.ListOptions{})
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (uc *ContainerUseCase) GetContainer(
	ctx context.Context,
	id string,
) (*domain.Container, error) {
	listOption := &container.ListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{Key: "id", Value: id}),
	}

	containers, err := uc.containerService.ListContainers(ctx, listOption)
	if err != nil {
		return nil, errors.ErrFailedToList
	}

	if len(containers) == 0 {
		return nil, errors.ErrContainerNotFound
	} else if len(containers) > 1 {
		return nil, errors.ErrMultipleContainers
	}

	return &containers[0], nil
}

func (uc *ContainerUseCase) StopContainer(ctx context.Context, id string) error {
	err := uc.containerService.StopContainer(ctx, id, &container.StopOptions{})
	if err != nil {
		return err
	}
	return nil
}
