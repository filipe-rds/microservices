# Build stage
FROM golang:1.24-alpine AS builder

# Install protoc and dependencies
RUN apk add --no-cache protobuf protobuf-dev

WORKDIR /app

# Copy proto files and generation script
COPY microservices-proto/ ../microservices-proto/

# Generate protobuf files using the flexible script
WORKDIR /microservices-proto
RUN chmod +x generate.sh && ./generate.sh shipping

# Return to app directory and copy shipping service files
WORKDIR /app
COPY microservices/shipping/go.mod microservices/shipping/go.sum ./

# Download dependencies
RUN go mod download

# Copy shipping source code
COPY microservices/shipping/ ./

# Update go.mod with new protobuf dependencies
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 3002

# Command to run
CMD ["./main"]
