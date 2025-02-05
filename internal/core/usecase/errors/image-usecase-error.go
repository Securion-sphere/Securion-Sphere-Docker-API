package errors

import "errors"

var (
	ErrImageNotFound     = errors.New("image not found")
	ErrImageLoadFailed   = errors.New("failed to load image")
	ErrMultipleImages    = errors.New("multiple image found, cannot proceed")
	ErrImageAlreadyExist = errors.New(
		"uploaded image is existed, please remove the old one and upload again",
	)
	ErrFailedToListAfterLoad = errors.New("failed to verify loaded, cannot proceed")
)
