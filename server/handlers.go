package main

import (
	"context"
	"log"

	"github.com/NearlyUnique/grpc-dojo/server/dojo"
)

type dojoServer struct{}

func (dojoServer) Hello(ctx context.Context, request *dojo.HelloRequest) (*dojo.HelloResponse, error) {
	log.Printf("received %q", request.Name)
	return &dojo.HelloResponse{
		Message: "Everything is ok, " + request.Name,
	}, nil
}
