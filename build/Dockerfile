FROM golang:1.21 AS builder

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/auth ./cmd/auth/main.go

FROM alpine:latest

WORKDIR /build

COPY configs /build/configs
COPY --from=builder /build/bin/auth /build/bin/auth

CMD ["./bin/auth"]
