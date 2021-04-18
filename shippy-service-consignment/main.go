// shippy/shippy-service-consignment/main.go

package main

import (
	"fmt"
	"log"

	// Import the generated protobuf code
	"context"

	pb "github.com/KapitanD/shipping-service/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/KapitanD/shipping-service/shippy-service-vessel/proto/vessel"
	micro "github.com/micro/go-micro/v2"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type vesselService struct {
	repo         repository
	vesselClient vesselProto.VesselService
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *vesselService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// Here we call a client instance of our vessel service with our consignment weight,
	// and the amount of containers as the capacity value
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id
	log.Printf("Create Consignment: %v", req)

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *vesselService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {

	repo := &Repository{}

	// Create a new service. Optionally include some options here.
	// Set-up micro instance
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)
	log.Println("Service Created")
	fmt.Println("Print smth")

	service.Init()

	vesselClient := vesselProto.NewVesselService("shippy.service.vessel",
		service.Client())

	log.Println("Connected to vessel service")

	// Register handlers
	if err := pb.RegisterShippingServiceHandler(service.Server(),
		&vesselService{repo, vesselClient}); err != nil {

		log.Panic(err)
	}

	log.Println("Register handlers")
	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
