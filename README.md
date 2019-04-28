# grpc-custom-error-sample

a template repository of grpc-go + grpc-gateway

## Features

* manage external proto files by [protodep](https://github.com/stormcat24/protodep)
* additional samples
  * [Customizing Error Response](https://github.com/soichisumi-sandbox/grpc-custom-error-sample)
  * [Handling request parameter in an interceptor by using **grpc-middleware/tags**](https://github.com/soichisumi-sandbox/grpc-middleware-sample)
  * [Authentication in an interceptor](https://github.com/soichisumi/grpc-auth-sample)

## Installation

* Install go 1.12 or above
  * brew install go
* Install [Protocol Buffers compiler](https://github.com/protocolbuffers/protobuf)
  * brew install protobuf
* Install [grpc-go](https://github.com/grpc/grpc-go)
  * go get -u google.golang.org/grpc
* Install [go protocol buffers plugin](https://github.com/golang/protobuf/tree/master/protoc-gen-go)
  * go get -u github.com/golang/protobuf/protoc-gen-go
* Install [grpc-gateway protocol buffers plugin](https://github.com/grpc-ecosystem/grpc-gateway)
  * go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway


## Usage

### Build

* `make proto-build`
* `make go-build`

### Run

* launch `cmd/api/api` and `cmd/gw/gw`

### Request example

* `curl localhost:8080/data?param=true` 
  
## Repository layout

* cmd
  * main applications
* proto
  * proto files
* app
  * internal packages