# Build stage (explicit amd64)
FROM --platform=linux/amd64 golang:1.21-alpine AS builder
WORKDIR /app

# Copy module files first for better layer caching
COPY client/go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build specifically for amd64 (redundant but explicit)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o /polygon-client .

# Final stage (amd64 only)
FROM --platform=linux/amd64 alpine:3.19
RUN apk --no-cache add ca-certificates

# Install binary at standard executable path
COPY --from=builder /polygon-client /usr/bin/polygon-client

# Verify architecture (optional but recommended)
RUN apk add --no-cache file && \
    file /usr/bin/polygon-client | grep "x86-64" && \
    apk del file

# Add non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

ENTRYPOINT ["polygon-client"]