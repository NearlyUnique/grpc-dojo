FROM grpc-build AS build-env

ADD ./*.go /go/src/github.com/NearlyUnique/grpc-dojo/server/
ADD ./dojo/*.go /go/src/github.com/NearlyUnique/grpc-dojo/server/dojo/

WORKDIR /go/src/github.com/NearlyUnique/grpc-dojo/server

RUN go build -o /app

# final stage
FROM alpine
COPY --from=build-env /app /app

WORKDIR /

ENTRYPOINT ["./app"]
