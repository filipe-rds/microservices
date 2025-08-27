package api

import (
	"context"
	"github.com/filipe-rds/microservices/shipping/internal/application/core/domain"
	"github.com/filipe-rds/microservices/shipping/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a *Application) CreateShipping(ctx context.Context, shipping domain.Shipping) (domain.Shipping, error) {
	err := a.db.Save(&shipping)
	if err != nil {
		return domain.Shipping{}, err
	}
	return shipping, nil
}
