module github.com/filipe-rds/microservices/order

go 1.24.5

require github/filipe-rds/microservices-proto/golang/order v0.0.0-00010101000000-000000000000
replace github/filipe-rds/microservices-proto/golang/order => ../../microservices-proto/golang/order

require github/filipe-rds/microservices-proto/golang/payment v0.0.0-00010101000000-000000000000
replace github/filipe-rds/microservices-proto/golang/payment => ../../microservices-proto/golang/payment