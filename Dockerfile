# Multi-stage Dockerfile for Crossplane Spy
# Builds both frontend (Next.js) and backend (Go) in a single image

# Stage 1: Build Frontend
FROM node:20-alpine AS frontend-builder

WORKDIR /build/frontend

# Copy frontend dependencies
COPY frontend/package*.json ./
RUN npm ci

# Copy frontend source
COPY frontend/ ./

# Build Next.js application in standalone mode
ENV NEXT_TELEMETRY_DISABLED=1
ENV NODE_ENV=production
RUN npm run build

# Stage 2: Build Backend
FROM golang:1.23-alpine AS backend-builder

WORKDIR /build/backend

# Install build dependencies
RUN apk add --no-cache git ca-certificates

# Copy Go modules files
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy backend source
COPY backend/ ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o crossplane-spy ./cmd/server

# Stage 3: Final Runtime Image
FROM alpine:3.19

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

WORKDIR /app

# Copy backend binary
COPY --from=backend-builder /build/backend/crossplane-spy ./crossplane-spy

# Copy frontend build output
COPY --from=frontend-builder /build/frontend/.next/standalone ./public/_next/standalone
COPY --from=frontend-builder /build/frontend/.next/static ./public/_next/static
COPY --from=frontend-builder /build/frontend/public ./public

# Change ownership
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Set environment variables
ENV PORT=8080
ENV GIN_MODE=release

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
CMD ["./crossplane-spy"]
