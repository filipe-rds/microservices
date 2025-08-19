package ports

import (
	"context"
	"github.com/filipe-rds/microservices/order/internal/application/core/domain"
)

type PaymentPort interface{
	Charge(ctx context.Context, order *domain.Order) error
}