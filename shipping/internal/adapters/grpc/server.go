package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/filipe-rds/microservices/shipping/internal/application/core/domain"
	"github.com/filipe-rds/microservices/shipping/internal/ports"
	shippingpb "github.com/filipe-rds/microservices-proto/golang/shipping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.APIPort
	port int
	shippingpb.UnimplementedShippingServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a *Adapter) Create(ctx context.Context, request *shippingpb.CreateShippingRequest) (*shippingpb.CreateShippingResponse, error) {
	var shippingItems []domain.ShippingItem
	for _, item := range request.Items {
		shippingItems = append(shippingItems, domain.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	newShipping := domain.NewShipping(request.OrderId, shippingItems)
	result, err := a.api.CreateShipping(ctx, newShipping)
	if err != nil {
		return nil, err
	}

	return &shippingpb.CreateShippingResponse{
		ShippingId:   result.ID,
		DeliveryDays: result.DeliveryDays,
	}, nil
}

func (a *Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	shippingpb.RegisterShippingServer(grpcServer, a)
	reflection.Register(grpcServer)

	log.Printf("starting gRPC server on port %d", a.port)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port")
	}
}
