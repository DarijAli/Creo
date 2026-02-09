# Build stage
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Copy only go.mod for dependency resolution
COPY . .

RUN go mod tidy
RUN go mod download



# Build the binary from main.go
RUN go build -o seed_register ./src/main

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/seed_register .

ENTRYPOINT ["./seed_register"]
