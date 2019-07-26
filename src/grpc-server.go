package main

import (
	"log"
	"net"
	"time"

	protodata "./protodata"

	"google.golang.org/grpc"
)

const (
	inputPort = ":5000"
)

type grpcServer struct{}

func (gs *grpcServer) Messaging(in *protodata.InfoRequest, stream protodata.ReplyStreamer_MessagingServer) error {
	log.Print("Get Info request...")
	log.Printf("Need to response Info with Name: %s Num: %d ", in.Mes.Name, in.Mes.Num)
	// First response
	if err := stream.Send(&protodata.InfoReply{Mes: &protodata.Info{Name: in.Mes.Name, Num: in.Mes.Num}}); err != nil {
		return err
	}
	time.Sleep(3 * time.Second)
	// Second response
	if err := stream.Send(&protodata.InfoReply{Mes: &protodata.Info{Name: in.Mes.Name, Num: 42}}); err != nil {
		return err
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", inputPort)

	if err != nil {
		log.Fatalf("gRPC server fails on listen: %v", err)
	}

	gs := grpc.NewServer()
	protodata.RegisterReplyStreamerServer(gs, &grpcServer{})

	log.Println("gRPC server registered")
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("Serve failed: %v", err)
	}
}
