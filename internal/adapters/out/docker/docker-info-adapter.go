package docker

import (
	"context"

	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/core/ports"
	"github.com/Securion-Sphere/Securion-Sphere-Docker-API/internal/infrastructure"
	"github.com/docker/docker/api/types/system"
)

type DockerInfoAdapter struct {
	dockerClient *infrastructure.DockerClient
}

var _ ports.InfoService = (*DockerInfoAdapter)(nil)

func NewDockerInfoAdapter(dc *infrastructure.DockerClient) *DockerInfoAdapter {
	return &DockerInfoAdapter{dockerClient: dc}
}

func (di *DockerInfoAdapter) GetInfo(ctx context.Context) (*system.Info, error) {
	resp, err := di.dockerClient.Client.Info(ctx)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
