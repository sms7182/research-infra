FROM golang:1.26.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/api-gateway ./cmd

FROM alpine:latest


RUN apk --no-cache add ca-certificates

WORKDIR /app


COPY --from=builder /app/api-gateway .

EXPOSE 8000
ENTRYPOINT ["./api-gateway"]
