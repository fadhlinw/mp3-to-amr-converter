# Base stage for ffmpeg
FROM jrottenberg/ffmpeg:4.4-alpine as ffmpeg_base

# Builder stage for Go application
FROM golang:1.20-alpine as builder

# Set working directory
WORKDIR /app

# Copy application files
COPY . .

# Install dependencies and build the Go application
RUN apk add --no-cache git \
  && go mod init convertapp || true \
  && go mod tidy \
  && go build -o convert .

# Final stage combining ffmpeg and Go binary
FROM jrottenberg/ffmpeg:4.4-alpine

# Set working directory
WORKDIR /app

# Copy Go binary from builder
COPY --from=builder /app/convert /app/convert

# Copy other necessary files
COPY . .

# Set default command
CMD ["./convert"]
