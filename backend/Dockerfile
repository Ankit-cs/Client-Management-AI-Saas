FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the main server API
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/api/main.go

# Build the database migrator
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/migrate ./cmd/migrate/main.go

# Final stage - using a minimal alpine image
FROM alpine:latest

WORKDIR /app

# Add required certificates for external HTTPS requests (like N8N & Google OAuth)
RUN apk --no-cache add ca-certificates tzdata

# Copy the built binaries from the builder stage
COPY --from=builder /app/server .
COPY --from=builder /app/migrate .

# Expose the API port
EXPOSE 8080

# Run the server by default
CMD ["./server"]
