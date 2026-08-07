FROM golang:1.26.0 as builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

FROM scratch
COPY --from=builder /app/main /app/main
ENTRYPOINT ["/app/main"]
