FROM grpc-build AS build-env

ADD ./*.go /go/src/github.com/NearlyUnique/grpc-dojo/client/
ADD ./dojo/*.go /go/src/github.com/NearlyUnique/grpc-dojo/client/dojo/

WORKDIR /go/src/github.com/NearlyUnique/grpc-dojo/client

RUN go build -o /app

# final stage
FROM alpine
COPY --from=build-env /app /app

WORKDIR /

ENTRYPOINT ["./app"]
