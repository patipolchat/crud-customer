# Stage 1: Build the Go application
FROM golang:1.23.0-alpine3.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o crud-customer

# Stage 2: Create a minimal image with Alpine to run the Go application
FROM alpine:3.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/crud-customer .

# Command to run the Go application
CMD ["./crud-customer", "serveApi"]