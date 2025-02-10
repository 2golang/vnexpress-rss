# ---- Build Stage ----
FROM golang:1.23 AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (improves caching)
COPY server/go.mod server/go.sum ./
COPY server/.env.sample ./.env
COPY server/Makefile ./Makefile

# Download dependencies
RUN make setup
RUN make install_packages

# Copy the entire project
COPY server/ .

# Build the Go binary
RUN make build

# Expose application port
EXPOSE 8080

# Run the application
CMD ["./main"]
