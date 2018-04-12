# Beyond HTTP+json

## Pre requisites

### Native tools

1. install go https://golang.org
1. install protoc https://github.com/google/protobuf/releases
1. install grpcurl `go get -u github.com/fullstorydev/grpcurl` cd into that folder then `go install ./...`

### Docker only

1. `docker pull znly/protoc`
1. copy [Dockerfile-build]
1. `docker build -t grpc-build -f Docker-build .`

### Optional Editor tooling

- VS Code https://code.visualstudio.com/ + extension "Go for Visual Studio Code"
- GoLand https://www.jetbrains.com/go/
