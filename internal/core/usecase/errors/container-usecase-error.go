package errors

import "errors"

var (
	ErrContainerNotFound       = errors.New("container not found")
	ErrContainerCreate         = errors.New("failed to create container")
	ErrContainerStart          = errors.New("failed to start container")
	ErrContainerExited         = errors.New("started container is exited")
	ErrMultipleContainers      = errors.New("multiple containers found, cannot proceed")
	ErrContainerStopAfterStart = errors.New(
		"failed to stop container after multiple containers started",
	)
	ErrContainerStop          = errors.New("failed to stop container")
	ErrContainerRemove        = errors.New("failed to remove container")
	ErrFailedToList           = errors.New("failed to list containers")
	ErrFailedToListAfterStart = errors.New("failed to verify started container, cannot proceed")
)
