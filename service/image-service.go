package service

import (
	"docker-api/domain"
	"docker-api/dto"
)

type ImageService struct{}

func NewImageService() domain.ImageService {
	return &ImageService{}
}

func (s *ImageService) GetImage() ([]domain.Image, error) { return []domain.Image{}, nil }
func (s *ImageService) LoadImage(uploadImageDto dto.UploadImageDto) (*domain.Image, error) {
	return &domain.Image{}, nil
}
func (s *ImageService) DeleteImage(deleteImageDto dto.DeleteImageDto) (map[string]string, error) {
	return map[string]string{"msg": "Deleted Image"}, nil
}
