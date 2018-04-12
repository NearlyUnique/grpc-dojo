//go:generate protoc --go_out=plugins=grpc:./dojo server.proto

package main

import (
	"context"
	"log"
	"os"

	"github.com/NearlyUnique/grpc-dojo/client/dojo"
	"google.golang.org/grpc"
)

func main() {
	host := ":9090"
	log.Printf("Running client on %q", host)

	// create a connection to our grpc server
	// as we have no certificates need grpc.WithInsecure()
	// let's encrypt is supported
	conn, err := grpc.Dial(host, grpc.WithInsecure())

	if err != nil {
		log.Printf("Failed to connect with %q, %v", host, err)
		os.Exit(1)
	}

	// client was generated for us by protoc
	client := dojo.NewDojoClient(conn)

	// make our call
	response, err := client.Hello(context.Background(), &dojo.HelloRequest{Name: "Punxsutawney Phil"})

	if err != nil {
		log.Printf("call failed: %v", err)
		os.Exit(1)
	}

	log.Printf("Got %v\n", *response)
}
