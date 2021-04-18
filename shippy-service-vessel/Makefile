build:
	protoc --proto_path=. --go_out=. --micro_out=. \
		proto/vessel/vessel.proto
	docker build -t shippy-service-vessel .

run: build
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 \
		shippy-service-vessel
