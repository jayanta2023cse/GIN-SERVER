# Build stage
FROM golang:1.23-bookworm AS builder

ENV GOPROXY=https://proxy.golang.org

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app source code
COPY . .

# Build the Go app binary statically (optional CGO disabled for Alpine)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app .

# Final stage
FROM alpine:latest

# Install certs if needed (for https/db connections)
RUN apk add --no-cache ca-certificates

WORKDIR /bin

# Copy built binary from builder
COPY --from=builder /bin/app .

# Copy environment files if you want to inject them into the container
COPY .env /bin/.env
COPY .env.prod /bin/.env.prod

ENV GO_ENV=prod

# Expose port your app listens on
EXPOSE 8080

# Run the app
CMD ["./app"]
