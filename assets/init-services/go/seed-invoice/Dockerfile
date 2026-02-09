# Build stage
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod download

# Build the binary from main.go
RUN go build -o seed_invoice ./src/main

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/seed_invoice .

ENTRYPOINT ["./seed_invoice"]
