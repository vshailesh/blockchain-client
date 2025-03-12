FROM --platform=linux/amd64 golang:1.21-alpine AS builder
WORKDIR /app

# Copy module files first for better layer caching
COPY client/ ./
RUN go mod download

# Copy source code
COPY . .

RUN cd client
RUN go build -o ../polygon-client

FROM --platform=linux/amd64 alpine:3.19
RUN apk --no-cache add ca-certificates

# Install binary at standard executable path
COPY --from=builder /polygon-client /usr/bin/polygon-client

# Add non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

ENTRYPOINT ["/bin/sh"]