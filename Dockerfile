FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /workspace

# Copy module definition files first
COPY go.mod go.sum ./

# Copy the rest of the source code
COPY cmd/ ./cmd/

# Build the Go Program
RUN go build -o /workspace/bin/client ./cmd/main.go

# Runtime image
FROM alpine:latest
WORKDIR /root
COPY --from=builder /workspace/bin/client ./client