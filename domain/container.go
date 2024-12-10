package domain

import (
	"github.com/docker/docker/api/types"

	"docker-api/dto"
)

type Container struct {
	ID      string
	Names   []string
	Image   string
	Ports   []types.Port
	Status  string
	Created int64
}

type ContainerService interface {
	GetContainer() ([]Container, error)
	CreateContainer(createContainerDto dto.CreateContainerDto) (*Container, error)
	DeleteContainer(deleteContainerDto dto.DeleteContainerDto) (map[string]string, error)
}
