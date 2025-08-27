package ports

import (
	"context"
	"github.com/filipe-rds/microservices/shipping/internal/application/core/domain"
)

type APIPort interface {
	CreateShipping(ctx context.Context, shipping domain.Shipping) (domain.Shipping, error)
}
