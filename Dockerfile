# BUILD STAGE
FROM golang:1.24.3 AS build

WORKDIR /app

# Copy go.mod and go.sum from api folder
COPY api/go.mod api/go.sum ./api/

# Download dependencies
WORKDIR /app/api
RUN go mod tidy

# Copy the entire api project into the container
COPY api/ ./ 

# Build the binary from the main package in cmd/server
RUN go build -o /app/server ./cmd/server

# FINAL IMAGE
FROM debian:bookworm-slim

WORKDIR /app

# Copy the binary from build stage
COPY --from=build /app/server /usr/local/bin/server

# Optional: expose port (replace with your actual port)
EXPOSE 8080

# Set the entry point
ENTRYPOINT ["server"]