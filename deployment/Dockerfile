FROM golang:1.24.1 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from a minimal base image (Alpine)
FROM alpine:latest

# Install libc6 compatibility (for Alpine)
RUN apk add --no-cache libc6-compat

WORKDIR /app

# Copy the compiled binary and required files
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Run the executable
RUN chmod +x /app/main
CMD ["./main"]
