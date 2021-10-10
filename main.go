package main

import (
	"github.com/Skiiru/goGrpcApi/pkg/api"
	"github.com/Skiiru/goGrpcApi/pkg/service"
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
