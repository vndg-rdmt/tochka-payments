FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/bin/app ./cmd/paymentspod/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/app .

CMD ["/app/app"]