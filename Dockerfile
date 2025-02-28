FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk add --no-cache git make

# Copy go.mod and go.sum files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN make build

# Create a minimal production image
FROM alpine:3.19

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy the binary from the builder stage
COPY --from=builder /app/bin/app /app/app

# Copy configuration files
COPY --from=builder /app/config /app/config

# Expose port
EXPOSE 8080

# Set the entry point
ENTRYPOINT ["/app/app"] 