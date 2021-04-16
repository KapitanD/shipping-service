package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pbf "github.com/KapitanD/shipping-service/consignment"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

// function parse gotten file
func parseFile(file string) (*pbf.Consignment, error) {
	var consignment *pbf.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, nil
}

func main() {
	// Estable connect to server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can not connect: %v", err)
	}
	defer conn.Close()
	client := pbf.NewShippingServiceClient(conn)

	// Pass consignment.json in processing,
	// else get path to file from commad line args
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Impossible to parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Can not create: %v", err)
	}

	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
