build:
	protoc --proto_path=. --go_out=. --micro_out=. \
		proto/vessel/vessel.proto
	docker build -t shippy-service-vessel .

run: build
	docker run \
		shippy-service-vessel
