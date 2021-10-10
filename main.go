package main

import (
	"goGrpcApi/pkg/api"
	"goGrpcApi/pkg/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	serverService := new(service.Server)
	api.RegisterApiServiceServer(grpcServer, serverService)
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer.Serve(listener)
}
