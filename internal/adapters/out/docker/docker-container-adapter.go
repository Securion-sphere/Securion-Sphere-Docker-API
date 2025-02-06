package docker

import (
	"context"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/domain"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/ports"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/infrastructure"
	"github.com/docker/docker/api/types/container"
)

type DockerContainerAdapter struct {
	dockerClient *infrastructure.DockerClient
}

// Ensure DockerContainerAdapter implements the ContainerService interface
var _ ports.ContainerService = (*DockerContainerAdapter)(nil)

// NewDockerContainerAdapter injects the shared Docker client
func NewDockerContainerAdapter(dc *infrastructure.DockerClient) *DockerContainerAdapter {
	return &DockerContainerAdapter{dockerClient: dc}
}

func (dc *DockerContainerAdapter) CreateContainer(
	ctx context.Context,
	config *container.Config,
	hostConfig *container.HostConfig,
) (string, error) {
	resp, err := dc.dockerClient.Client.ContainerCreate(ctx, config, hostConfig, nil, nil, "")
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func (dc *DockerContainerAdapter) ListContainers(
	ctx context.Context,
	listOption *container.ListOptions,
) ([]domain.Container, error) {
	containers, err := dc.dockerClient.Client.ContainerList(ctx, *listOption)
	if err != nil {
		return nil, err
	}

	var result []domain.Container
	for _, c := range containers {
		result = append(result, domain.Container{ID: c.ID, Name: c.Names[0], Image: c.ID, Ports: c.Ports, Status: c.State, Created: c.Created})
	}
	return result, nil
}

func (dc *DockerContainerAdapter) StartContainer(
	ctx context.Context,
	id string,
	startOptions *container.StartOptions,
) error {
	return dc.dockerClient.Client.ContainerStart(ctx, id, *startOptions)
}

func (dc *DockerContainerAdapter) StopContainer(
	ctx context.Context,
	id string,
	stopOptions *container.StopOptions,
) error {
	return dc.dockerClient.Client.ContainerStop(ctx, id, *stopOptions)
}

func (dc *DockerContainerAdapter) RemoveContainer(
	ctx context.Context,
	id string,
	removeOptions *container.RemoveOptions,
) error {
	return dc.dockerClient.Client.ContainerRemove(ctx, id, *removeOptions)
}
