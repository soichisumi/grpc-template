proto-build:
	protoc -I=./proto -I=./proto/ext --go_out=plugins=grpc:./app/proto --grpc-gateway_out=./app/proto proto/*.proto

go-build:
	GO111MODULE=on go build -o cmd/api/api ./cmd/api/
	GO111MODULE=on go build -o cmd/gw/gw ./cmd/gw/

protodep:
	protodep up