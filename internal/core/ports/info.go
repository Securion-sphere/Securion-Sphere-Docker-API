package ports

import (
	"context"

	"github.com/docker/docker/api/types/system"
)

type InfoService interface {
	GetInfo(ctx context.Context) (*system.Info, error)
}
