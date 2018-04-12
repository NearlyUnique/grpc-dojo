# Build a server

## Native

1. `go generate`
1. `go build && server`

## Docker

1. _go generate_
    - Linux `docker run --rm -v $(pwd):/app -w=/app znly/protoc --go_out=plugins=grpc:./dojo -I. server.proto`
    - Windows `docker run --rm -v %CD:\=/%:/app -w=/app znly/protoc --go_out=plugins=grpc:./dojo -I. server.proto`
1. `docker build -t grpc02 .`
1. `docker run --rm grpc02`

## grpcurl

### Generate the protoset files

This describes the service (there are also ways to get the service to describe itself)

```bash
protoc --proto_path=. --descriptor_set_out=server.protoset ./server.proto
```

### Make a request

```bash
cat helloRequest.json | grpcurl -plaintext -d @ -protoset server.protoset localhost:9090 dojo.Dojo/Hello
```