package main

import (
	"log"

	"github.com/filipe-rds/microservices/shipping/config"
	"github.com/filipe-rds/microservices/shipping/internal/adapters/db"
	"github.com/filipe-rds/microservices/shipping/internal/adapters/grpc"
	"github.com/filipe-rds/microservices/shipping/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURl())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
