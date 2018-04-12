# Beyond HTTP+json

## Pre requisites

### Native tools

1. install go https://golang.org
1. install protoc https://github.com/google/protobuf/releases
1. install grpcurl `go get -u github.com/fullstorydev/grpcurl`
    1. find your GOPATH, `go env GOPATH`
    1. `cd ${GOPATH}` or `cd %GOPATH%`
    1. `go install ./...`

### Docker only

1. `docker pull znly/protoc`
1. copy [Dockerfile-build]
1. `docker build -f Dockerfile-build -t grpc-build .`

### Optional Editor tooling

- VS Code https://code.visualstudio.com/ + extension "Go for Visual Studio Code"
- GoLand https://www.jetbrains.com/go/

## Running order

- [[check-tools/readme.md]]
- [[server/readme.md]]
- [[client/readme.md]]