# Build a server

1. make go generate work
1. connect in grpc
1. register service
1. service type
1. implement interface
1. test with grpcurl

## Native

1. `go generate`
1. build and run
    - Windows `go build -o app.exe && app`
    - Linus `go build -o app && app`

## Docker

1. _go generate_
    - Linux `docker run --rm -v $(pwd):/app -w=/app znly/protoc --go_out=plugins=grpc:. -I. server.proto`
    - Windows `docker run --rm -v %CD:\=/%:/app -w=/app znly/protoc --go_out=plugins=grpc:. -I. server.proto`
1. `docker build -t grpc02 .`
1. `docker run --rm grpc02`