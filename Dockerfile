FROM golang:1.20.3-alpine AS builder

COPY . /github.com/armanbektassov/go_auth/grpc/source/
WORKDIR /github.com/armanbektassov/go_auth/grpc/source/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/armanbektassov/go_auth/grpc/source/bin/auth_server .

CMD ["./auth_server"]