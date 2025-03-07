# Use official Go image as the base for building
FROM golang:1.24-alpine AS builder

# Set working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application with static linking for a smaller image
RUN CGO_ENABLED=0 GOOS=linux go build -o blockchain-client .

FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/blockchain-client .

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./blockchain-client"]
