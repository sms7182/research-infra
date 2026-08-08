FROM golang:1.26.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/auth-service ./cmd

FROM alpine:latest


RUN apk --no-cache add ca-certificates

WORKDIR /app


COPY --from=builder /app/auth-service .

EXPOSE 8082
ENTRYPOINT ["./auth-service"]
