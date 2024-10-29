# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files have not changed
RUN go mod download

# Copy the entire project directory to the container
COPY . .

# Build the Go app, output as a binary called "app"
RUN go build -o app .

# Stage 2: Create a minimal runtime image to run the binary
FROM alpine:latest

# Set the working directory inside the runtime container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the binary
CMD ["./app"]