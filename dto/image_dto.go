package dto

import "mime/multipart"

type UploadImageDto struct {
	File multipart.File
}

type DeleteImageDto struct {
	Name string `json:"name" validate:"required"`
}
