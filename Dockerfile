# Use the official Golang image for building the application
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the entire application code
COPY . .

# Build the application
RUN go build -o server ./cmd/server

# Use a minimal image for the final stage
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./server"]