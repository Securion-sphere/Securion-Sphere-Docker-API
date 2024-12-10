package domain

import (
	"docker-api/dto"
)

type Image struct {
	ID   string
	Name string
	Size string
}

type ImageService interface {
	GetImage() ([]Image, error)
	LoadImage(loadImageDto dto.UploadImageDto) (*Image, error)
	DeleteImage(deleteImageDto dto.DeleteImageDto) (map[string]string, error)
}
