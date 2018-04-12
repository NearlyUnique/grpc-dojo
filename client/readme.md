# Build a client

## Native

1. `go generate`
1. `go build && client`

## Docker

1. _go generate_
    - Linux `docker run --rm -v $(pwd):/app -w=/app znly/protoc --go_out=plugins=grpc:. -I. server.proto`
    - Windows `docker run --rm -v %CD:\=/%:/app -w=/app znly/protoc --go_out=plugins=grpc:. -I. server.proto`
1. `docker build -t grpc02 .`
1. `docker run --rm grpc02`