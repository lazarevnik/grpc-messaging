# Simple project implements some features testing gRPC capabilities with Go lang

## Installed software
This project uses 
- Go 1.12.7 
- [protoc](https://github.com/protocolbuffers/protobuf/blob/master/src/README.md) with [Go support](https://github.com/golang/protobuf)

Check PATH and GOPATH env to correct work.

## gRPC code generator
To run Go generator use such command:

`protoc -I src/ src/protodata/protofile.proto --go_out=plugins=grpc:./src`

After this you can find Generated types and methods for gRPC in src/protodata/protofile.pb.go 

## Running
To test notifications
1. Run grpc server:

  `go run src/grpc-server.go`
  
2. Run grpc client:

  `go run src/grpc-client.go`
  
3. Run command:

  `curl localhost:8083/test`
