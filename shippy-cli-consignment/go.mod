module github.com/KapitanD/shipping-service/shippy-cli-consignment

replace github.com/KapitanD/shipping-service/shippy-service-consignment => ../shippy-service-consignment

go 1.16

require (
	github.com/KapitanD/shipping-service/shippy-service-consignment v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.37.0
)
