package domain

import "github.com/docker/docker/api/types"

type Container struct {
	ID      string
	Name    string
	Image   string
	Ports   []types.Port
	Status  string
	Created int64
}
