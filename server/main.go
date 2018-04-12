//go:generate protoc --go_out=plugins=grpc:./dojo server.proto

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/NearlyUnique/grpc-dojo/server/dojo"
	"google.golang.org/grpc"
)

func main() {
	host := ":9090"
	log.Printf("Running server on %q", host)

	// 1. listen on a port
	lis, err := net.Listen("tcp", host)

	if err != nil {
		log.Printf("Cannot start server %v", err)
	}

	// 2. create a grpc server
	svr := grpc.NewServer()

	// 3. tell grpc about OUR handler
	dojo.RegisterDojoServer(svr, dojoServer{})

	// start a go routine (light weight thread)
	go func() {
		// tell the grpc server (with our registered handlers) to listen the tcp connection
		if err := svr.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	fmt.Println("CTRL+C to quit")
	waitForShutdown()
	log.Printf("Stopping...")
	svr.Stop()
}

func waitForShutdown() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
}
