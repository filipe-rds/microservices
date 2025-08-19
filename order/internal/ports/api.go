package ports

import (
	"context"
	"github.com/filipe-rds/microservices/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(ctx context.Context,order domain.Order) (domain.Order, error)
}
