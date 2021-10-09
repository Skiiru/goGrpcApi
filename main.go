package main

import (
	"google.golang.org/grpc"
	"grpcApi/pkg/api"
	"grpcApi/pkg/service"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	serverService := new(service.Server)
	api.RegisterApiServiceServer(grpcServer, serverService)
	listener, error := net.Listen("tcp", "localhost:8081")
	if error != nil {
		log.Fatalf("failed to listen: %v", error)
	}
	grpcServer.Serve(listener)
}
