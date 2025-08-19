package api

import (
	"context"
	"log"
	"time"
	"github.com/filipe-rds/microservices/order/internal/application/core/domain"
	"github.com/filipe-rds/microservices/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db: db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(ctx context.Context,order domain.Order) (domain.Order, error) {
	ctxTimeout , cancel := context.WithTimeout (context.Background(), 2*time.Second )
	defer cancel()
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}

	Quantity := int32(0)
	for _,item := range order.OrderItems {
		Quantity += item.Quantity
	}
	if Quantity > 50 {
		order.Status = "Canceled"
		quantityErr := a.db.Save(&order)
		if quantityErr != nil {
		return domain.Order{}, quantityErr
		}
		return order, status.Error(codes.InvalidArgument, "Quantity cannot be more than 50.")
	}

	paymentErr := a.payment.Charge(ctxTimeout, &order)
	if paymentErr != nil {
		if status.Code(paymentErr) == codes.DeadlineExceeded {
			log.Fatalf("Erro: %v",paymentErr)
		}
		order.Status = "Canceled"
		if saveErr := a.db.Save(&order); saveErr != nil {
			return domain.Order{}, saveErr
		}
		return order, paymentErr
	}

	order.Status = "Paid"
	updateErr := a.db.Save(&order)
	if updateErr != nil {
		return domain.Order{}, err
	}
	return order, nil
}

