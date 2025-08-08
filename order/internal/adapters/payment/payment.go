package payment_adapter

import (
	"context"
	"log"

	paymentpb "github.com/filipe-rds/microservices-proto/golang/payment"
	"github.com/filipe-rds/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type Adapter struct {
	payment paymentpb.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := paymentpb.NewPaymentClient(conn)

	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(), &paymentpb.CreatePaymentRequest{
		CostumerID: order.CostumerID,
		OrderId: order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}

