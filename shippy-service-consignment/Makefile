build:
	protoc --proto_path=. --go_out=. --micro_out=. \
		proto/consignment/consignment.proto
	docker build -t shippy-service-consigment .

run: build
	docker run -p 50051:50051 \
        -e MICRO_SERVER_ADDRESS=:50051 \
        shippy-service-consignment