package main

import (
	"log"

	"github.com/NearlyUnique/grpc-dojo/duplex/dojo"
)

type dojoServer struct{}

func (dojoServer) Hello(stream dojo.Dojo_HelloServer) error {
	log.Println("duplex request")
	return nil
}
