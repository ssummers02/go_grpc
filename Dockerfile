FROM golang:1.16-alpine3.12 as builder
COPY go.mod go.sum /go/src/go_grpc/
WORKDIR /go/src/go_grpc/
RUN go mod download
COPY . /go/src/go_grpc/
RUN go build -o build/go_grpc go_grpc

FROM alpine:latest
COPY --from=builder /go/src/go_grpc/build/go_grpc /usr/bin/app
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/app"]


