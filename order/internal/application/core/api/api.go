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
	db      ports.DBPort
	payment ports.PaymentPort
	shipping ports.ShippingPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort, shipping ports.ShippingPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
		shipping: shipping,
	}
}

func (a Application) PlaceOrder(ctx context.Context,order domain.Order) (domain.Order, error) {
	ctxTimeout, cancel := context.WithTimeout (context.Background(), 2*time.Second )
	defer cancel()
	
	// Validar se os produtos existem no banco
	var productCodes []string
	for _, item := range order.OrderItems {
		productCodes = append(productCodes, item.ProductCode)
	}
	
	err := a.db.ValidateProducts(productCodes)
	if err != nil {
		return domain.Order{}, status.Error(codes.InvalidArgument, err.Error())
	}
	
	err = a.db.Save(&order)
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
		return domain.Order{}, updateErr
	}

	// Chamar shipping apenas se o pagamento foi bem-sucedido
	_, shippingErr := a.shipping.CreateShipping(ctxTimeout, &order)
	if shippingErr != nil {
		log.Printf("Warning: Failed to create shipping for order %d: %v", order.ID, shippingErr)
		// Não falha o pedido se o shipping falhar, apenas loga o erro
	}

	return order, nil
}

