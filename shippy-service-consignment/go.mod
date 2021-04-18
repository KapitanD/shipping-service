module github.com/KapitanD/shipping-service/shippy-service-consignment

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/KapitanD/shipping-service/shippy-service-vessel v0.0.0-20210418112201-3a095f1d0df5
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/protobuf v1.26.0
)
