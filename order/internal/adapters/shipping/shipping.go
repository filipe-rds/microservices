package shipping

import (
	"context"
	"github.com/filipe-rds/microservices/order/internal/application/core/domain"
	shippingpb "github.com/filipe-rds/microservices-proto/golang/shipping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	shippingServiceUrl string
}

func NewAdapter(shippingServiceUrl string) (*Adapter, error) {
	return &Adapter{
		shippingServiceUrl: shippingServiceUrl,
	}, nil
}

func (a *Adapter) CreateShipping(ctx context.Context, order *domain.Order) (int32, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(a.shippingServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	shippingService := shippingpb.NewShippingClient(conn)

	var shippingItems []*shippingpb.ShippingItem
	for _, item := range order.OrderItems {
		shippingItems = append(shippingItems, &shippingpb.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	shippingRequest := &shippingpb.CreateShippingRequest{
		OrderId: order.ID,
		Items:   shippingItems,
	}

	shippingResponse, err := shippingService.Create(ctx, shippingRequest)
	if err != nil {
		return 0, err
	}

	return shippingResponse.DeliveryDays, nil
}
