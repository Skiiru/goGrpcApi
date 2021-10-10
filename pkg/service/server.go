package service

import (
	"context"
	"goGrpcApi/pkg/api"
)

type Server struct {
	api.UnimplementedApiServiceServer
}

func (s Server) Method(_ context.Context, request *api.MethodRequest) (*api.MethodResponse, error) {
	return &api.MethodResponse{Result: "Hello, " + request.Message}, nil
}
