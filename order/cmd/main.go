package cmd

import(
	"log"
	"github.com/filipe-rds/microservices/order/config"
	"github.com/filipe-rds/microservices/order/internal/adapters/db"
	"github.com/filipe-rds/microservices/order/internal/adapters/grpc"
	// "github.com/filipe-rds/microservices/order/internal/adapters/rest"
	"github.com/filipe-rds/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSouceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}