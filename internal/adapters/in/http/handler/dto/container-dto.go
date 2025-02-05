package dto

type CreateContainerDto struct {
	Image         string `json:"image"          validate:"required"`
	ContainerPort uint16 `json:"container_port" validate:"required"`
	HostPort      uint16 `json:"host_port"      validate:"required,min=40000,max=60000"`
	Flag          string `json:"flag"           validate:"required"`
}
