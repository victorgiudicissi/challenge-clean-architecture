FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o orders ./cmd/ordersystem

FROM alpine:latest AS runner

WORKDIR /root/

COPY --from=builder /app/orders ./
COPY .env ./

RUN chmod +x ./orders

EXPOSE 8000 8080 50051

CMD ["./orders"]
