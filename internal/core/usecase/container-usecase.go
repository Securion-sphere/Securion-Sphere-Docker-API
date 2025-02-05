package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/domain"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/ports"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/go-connections/nat"
	"github.com/labstack/echo/v4"
)

type ContainerUseCase struct {
	containerService ports.ContainerService
}

func NewContainerUseCase(containerService ports.ContainerService) *ContainerUseCase {
	return &ContainerUseCase{containerService: containerService}
}

func (uc *ContainerUseCase) CreateContainer(ctx context.Context, image string, containerPort uint16, hostPort uint16, flag string) (*domain.Container, error) {
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
		return nil, echo.NewHTTPError(http.StatusServiceUnavailable, err)
	}

	err = uc.containerService.StartContainer(ctx, id, &container.StartOptions{})
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusServiceUnavailable, err)
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
	if len(containers) > 1 {
		for _, c := range containers {
			if err := uc.containerService.RemoveContainer(ctx, c.ID, &container.RemoveOptions{}); err != nil {
				return nil, echo.NewHTTPError(http.StatusConflict, "Start the container successfully but multiple container id started and cannot be removed")
			}
		}
		return nil, err
	} else if len(containers) == 0 {
		return nil, echo.NewHTTPError(http.StatusConflict, "Failed to start container")
	} else if err != nil {
		uc.containerService.RemoveContainer(ctx, containers[0].ID, &container.RemoveOptions{})
		return nil, err
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
		return nil, echo.NewHTTPError(http.StatusServiceUnavailable, err)
	}
	return containers, nil
}

func (uc *ContainerUseCase) GetContainer(ctx context.Context, id string) (*domain.Container, error) {
	listOption := &container.ListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{Key: "id", Value: id}),
	}

	containers, err := uc.containerService.ListContainers(ctx, listOption)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusServiceUnavailable, err)
	}

	if len(containers) == 0 {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Container not found")
	} else if len(containers) > 1 {
		return nil, echo.NewHTTPError(http.StatusConflict, "More than one container found, try to specific more ID length")
	}

	return &containers[0], nil
}

func (uc *ContainerUseCase) StopContainer(ctx context.Context, id string) error {
	err := uc.containerService.StopContainer(ctx, id, &container.StopOptions{})
	if err != nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, err)
	}
	return nil
}
