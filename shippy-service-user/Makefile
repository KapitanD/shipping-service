.PHONY: proto
proto:
	protoc --proto_path=. --go_out=. --micro_out=. \
		proto/user/user.proto

build: proto
	docker build -t shippy-service-vessel .

run: build
	docker run \
		shippy-service-vessel
