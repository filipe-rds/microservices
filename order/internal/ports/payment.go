package ports

import "github.com/filipe-rds/microservices/order/internal/application/core/domain"

type PaymentPort interface{
	Charge(domain.Order) error
}