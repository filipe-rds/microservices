package api

import (
	"context"
	"log"

	"github.com/filipe-rds/microservices/payment/internal/application/core/domain"
	"github.com/filipe-rds/microservices/payment/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	log.Println("Iniciando o Charge...")
	err := a.db.Save(ctx, &payment)
	if err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}
