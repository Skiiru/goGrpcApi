package service

import (
	"context"
	"grpcApi/api"
)

type Server struct {
	api.UnimplementedApiServiceServer
}

func (s Server) Method(context context.Context, request *api.MethodRequest) (*api.MethodResponse, error) {
	return &api.MethodResponse{Result: "Hello world"}, nil
}
