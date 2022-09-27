FROM golang:alpine3.15 AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o wormhole -ldflags="-s -w" main.go

FROM alpine:3
WORKDIR /cmd
COPY --from=builder /build/wormhole ./

EXPOSE 8080
ENTRYPOINT ["./wormhole", "server", "--port", ":8080"]