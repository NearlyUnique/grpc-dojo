//go:generate protoc --go_out=plugins=grpc:. server.proto
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	m := HelloRequest{
		Name: "Hello World",
		List: []*SubType{
			&SubType{
				Enabled: true,
				Number:  42,                          // 0x2a
				Lookup:  map[string]int32{"one": 99}, // 0x63
			},
		},
	}

	jsonBody, _ := json.Marshal(m)
	fmt.Printf("json: %d bytes\n%s\n\n", len(jsonBody), string(jsonBody))
	protoBufBody, _ := proto.Marshal(&m)
	fmt.Printf("proto: %d bytes\n", len(protoBufBody))
	fmt.Print(hex.Dump(protoBufBody))
}
