package infrastructure

import (
	"github.com/docker/docker/client"
)

// DockerClient provides a shared Docker SDK client
type DockerClient struct {
	Client *client.Client
}

// NewDockerClient initializes a single Docker client instance
func NewDockerClient() (*DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &DockerClient{Client: cli}, nil
}
