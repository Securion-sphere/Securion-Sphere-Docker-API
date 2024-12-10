// service/container_service.go
package service

import (
	"context"
	"docker-api/domain"
	"docker-api/dto"
	"errors"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type ContainerService struct{}

func NewContainerService() domain.ContainerService {
	return &ContainerService{}
}

func (s *ContainerService) GetContainer() ([]domain.Container, error) {
	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	// List running containers
	containers, err := dockerClient.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	dockerClient.Close()

	// Prepare response data
	var containerInfos []domain.Container
	for _, container := range containers {
		containerInfos = append(containerInfos, domain.Container{
			ID:      container.ID,
			Names:   container.Names,
			Image:   container.Image,
			Ports:   container.Ports,
			Status:  container.Status,
			Created: container.Created,
		})
	}

	return containerInfos, nil
}

func (s *ContainerService) CreateContainer(createContainerDto dto.CreateContainerDto) (*domain.Container, error) {
	ctx := context.Background()

	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	containerPort := nat.Port(fmt.Sprintf("%d", createContainerDto.ContainerPort))

	portBindings := nat.PortMap{
		containerPort: []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: fmt.Sprintf("%d", createContainerDto.HostPort),
			},
		},
	}

	resp, err := dockerClient.ContainerCreate(ctx, &container.Config{
		Image: createContainerDto.Image,
		Env:   []string{"SC_FLAG=" + createContainerDto.Flag},
		ExposedPorts: nat.PortSet{
			containerPort: struct{}{},
		},
	}, &container.HostConfig{
		PortBindings: portBindings,
		AutoRemove:   true,
	}, nil, nil, "")
	if err != nil {
		return nil, err
	}

	err = dockerClient.ContainerStart(ctx, resp.ID, container.StartOptions{})
	if err != nil {
		return nil, err
	}

	filterArgs := filters.NewArgs()
	filterArgs.Add("id", resp.ID)

	containers, err := dockerClient.ContainerList(ctx, container.ListOptions{
		All:     true,
		Filters: filterArgs,
	})
	if err != nil {
		return nil, err
	} else if len(containers) == 0 {
		return nil, errors.New("created container ID not found")
	}
	container := containers[0]

	dockerClient.Close()

	return &domain.Container{
		ID:      container.ID,
		Names:   container.Names,
		Image:   container.Image,
		Ports:   container.Ports,
		Status:  container.Status,
		Created: container.Created,
	}, nil
}

func (s *ContainerService) DeleteContainer(deleteContainerDto dto.DeleteContainerDto) (map[string]string, error) {
	// Create a new Docker client
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	err = cli.ContainerStop(ctx, deleteContainerDto.ContainerId, container.StopOptions{})
	if err != nil {
		return nil, err
	}

	return map[string]string{"msg": "Deleted container successfully"}, nil
}
