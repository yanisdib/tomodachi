# Step 1: Build the Go application

# Use the official Golang image as base for building the Go app
FROM golang:1.23.4 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files from the server directory
COPY server/go.mod server/go.sum ./

# Run go mod tidy to download the dependencies and create the go.sum file
RUN go mod tidy

# Copy the rest of the Go project files (from server directory)
COPY server/ ./

# Build the Go binary (assuming the entry point is in the server folder)
RUN go build -o api .

# Step 2: Create a minimal production image

# Use a smaller base image for the final image (debian-slim in this case)
FROM debian:bookworm-slim

# Set the working directory for the production container
WORKDIR /app

# Copy the built Go binary from the build stage
COPY --from=build /app/api /usr/local/bin/api

# Set the entry point for the container to run the Go binary
ENTRYPOINT ["api"]

# Expose the port your application is running on (assuming it's 8080)
EXPOSE 8080
