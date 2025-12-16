FROM golang:latest AS builder

WORKDIR /app

COPY . .
COPY go.mod go.sum ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o subscription_service ./cmd/main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/subscription_service .
COPY internal/migration/sql/ internal/migration/sql/

EXPOSE 8080

CMD ["./subscription_service"]