package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	host := ":9090"
	log.Printf("Running server on %q", host)
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Printf("Cannot start server %v", err)
	}

	fmt.Printf("now connect to grpc %v", lis)

	fmt.Println("CTRL+C to quit")
	waitForShutdown()
}

func waitForShutdown() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
}
