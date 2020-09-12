package main

import (
	"log"
	"net"

	"github.com/kywang1/grpc_tutorial/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to liste on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	// Pass in grpcServer that we want to register this over
	// For this we want to register the chat service with our grpcServer
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
