# Start with a lightweight Go image
FROM golang:1.22 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o my-go-api

# Use a base image with a compatible glibc version
FROM debian:bookworm-slim

# Install necessary runtime dependencies
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the binary from the builder stage
COPY --from=builder /app/my-go-api /my-go-api

# Set up the application working directory
WORKDIR /app

# Expose the application port
EXPOSE 3000

# Run the application
CMD ["/my-go-api"]
