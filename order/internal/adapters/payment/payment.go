package payment_adapter

import (
	"context"
	"time"

	"github.com/filipe-rds/microservices-proto/golang/payment"
	"github.com/filipe-rds/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
    "google.golang.org/grpc/codes"
)


type Adapter struct {
	payment payment.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithCodes(codes.DeadlineExceeded,codes.ResourceExhausted, codes.Unavailable),
		grpc_retry.WithMax(5),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
		)))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// conn, err := grpc.Dial(paymentServiceUrl, opts...)
	conn, err := grpc.NewClient(paymentServiceUrl, opts...)

	if err != nil {
		return nil, err 
	}
	client := payment.NewPaymentClient(conn)
	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(), &payment.CreatePaymentRequest{
		CustomerId: order.CustomerId,
		OrderId: order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}

