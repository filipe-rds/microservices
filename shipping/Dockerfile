# Build stage
FROM golang:1.24-alpine AS builder

# Install protoc, git, bash and dependencies
RUN apk add --no-cache protobuf protobuf-dev git bash

WORKDIR /app

# Clone the microservices-proto repository
RUN git clone https://github.com/filipe-rds/microservices-proto.git ../microservices-proto

# Generate protobuf files using the script from the cloned repository
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
