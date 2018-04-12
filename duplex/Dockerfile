FROM grpc-build AS build-env

ADD ./*.go /go/src/github.com/NearlyUnique/grpc-dojo/duplex/
ADD ./dojo/*.go /go/src/github.com/NearlyUnique/grpc-dojo/duplex/dojo/

WORKDIR /go/src/github.com/NearlyUnique/grpc-dojo/duplex

RUN go build -o /app

# final stage
FROM alpine
COPY --from=build-env /app /app

WORKDIR /

ENTRYPOINT ["./app"]
