package main

import (
	"fmt"
	"log"
)

func main() {
	host := ":9090"
	log.Printf("Running client on %q", host)

	fmt.Printf("Got %v\n", *response)
}
