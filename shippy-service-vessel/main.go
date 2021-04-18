// shippy-service-vessel/main.go
package main

import (
	"context"
	"log"
	"os"

	pb "github.com/KapitanD/shipping-service/shippy-service-vessel/proto/vessel"
	micro "github.com/micro/go-micro/v2"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	service := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shippy").Collection("vessels")

	repo := &MongoRepository{vesselCollection}

	h := &handler{repo}

	// Register our implementation with
	if err := pb.RegisterVesselServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
