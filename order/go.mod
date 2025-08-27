module github.com/filipe-rds/microservices/order

go 1.24.5

require (
	github.com/filipe-rds/microservices-proto/golang/order v0.0.0-00010101000000-000000000000
	github.com/filipe-rds/microservices-proto/golang/payment v0.0.0-00010101000000-000000000000
	github.com/filipe-rds/microservices-proto/golang/shipping v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	google.golang.org/grpc v1.74.2
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.30.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.43.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
	google.golang.org/protobuf v1.36.8 // indirect
)

replace github.com/filipe-rds/microservices-proto/golang/order => ../../microservices-proto/golang/order
replace github.com/filipe-rds/microservices-proto/golang/payment => ../../microservices-proto/golang/payment
replace github.com/filipe-rds/microservices-proto/golang/shipping => ../../microservices-proto/golang/shipping
