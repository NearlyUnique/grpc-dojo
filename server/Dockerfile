FROM grpc-build AS build-env

ADD ./*.go /go/src/github.com/dojo/app/

WORKDIR /go/src/github.com/dojo/app

RUN go build -o /app

# final stage
FROM alpine
COPY --from=build-env /app /app

WORKDIR /

ENTRYPOINT ["./app"]
