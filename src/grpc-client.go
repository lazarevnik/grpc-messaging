package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"

	protodata "./protodata"
)

const (
	serverAddress = "localhost:5000"
	inputPort     = ":8083"
)

type grpcClient struct {
	client protodata.NotifierClient
}

func (gc *grpcClient) getConnection(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("gRPC connection error: %v", err)
	}

	gc.client = protodata.NewNotifierClient(conn)
}

func (gc grpcClient) sendToServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var infoName string = "Hello!"
	var infoNum int32 = 3

	stream, err := gc.client.Messaging(ctx, &protodata.InfoRequest{Mes: &protodata.Info{Name: infoName, Num: infoNum}})

	if err != nil {
		log.Fatalf("gRPC client failed when get stream: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("gRPC client Error while recieving: %v", err)
		}

		log.Println("Recieve : %s", message)

	}
	log.Println("Massaging stopped")
}

func (s grpcClient) handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler OK")
	s.sendToServer()
}

func main() {
	cli := grpcClient{}
	cli.getConnection(serverAddress)

	http.HandleFunc("/test", cli.handler)
	log.Print("Server started")
	log.Fatal(http.ListenAndServe(inputPort, nil))
}
