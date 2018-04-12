package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/NearlyUnique/grpc-dojo/duplex/dojo"
	"google.golang.org/grpc/metadata"
)

type dojoServer struct{}

func (dojoServer) Hello(stream dojo.Dojo_HelloServer) error {
	ch := make(chan string)
	ctx := stream.Context()
	_, traceID := extractMeta(ctx)
	var wg sync.WaitGroup
	wg.Add(2)

	// inbound
	go func() {
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				log.Printf("[%s] Closed", traceID)
				break
			}
			if err != nil {
				log.Printf("[%s] Recv failed: %v", traceID, err)
				break
			}

			ch <- req.Name
		}
		wg.Done()
	}()

	// out bound
	go func() {
		for {
			select {
			case name := <-ch:
				stream.Send(&dojo.HelloResponse{
					Message: fmt.Sprintf("Responding %s %v", name, time.Now().Format(time.Kitchen)),
				})
			case <-ctx.Done():
				log.Printf("[%s] Closing by force", traceID)
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()
	return nil
}

func extractMeta(ctx context.Context) (userAgent, traceID string) {
	if m, ok := metadata.FromIncomingContext(ctx); ok {
		if u, ok := m["user-agent"]; ok {
			userAgent = strings.Join(u, ";")
		}
		if u, ok := m["traceid"]; ok && len(u) == 1 {
			traceID = u[0]
		}
	}
	return userAgent, traceID
}
