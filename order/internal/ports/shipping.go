package ports

import (
	"context"
	"github.com/filipe-rds/microservices/order/internal/application/core/domain"
)

type ShippingPort interface{
	CreateShipping(ctx context.Context, order *domain.Order) (int32, error)
}
