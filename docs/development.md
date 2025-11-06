# Development Guide

## Prerequisites

### Required
- **Go**: 1.23 or later
- **Node.js**: 20 or later
- **Kubernetes cluster**: With Crossplane installed
- **kubectl**: Configured with cluster access

### Recommended
- **Docker**: For containerization
- **Helm**: For deployment
- **make**: For running common tasks

## Project Structure

```
crossplane-spy/
├── backend/                 # Go backend
│   ├── cmd/
│   │   └── server/         # Main application
│   ├── internal/
│   │   ├── api/            # REST API handlers
│   │   ├── k8s/            # Kubernetes client
│   │   └── models/         # Data models
│   └── pkg/                # Public packages
├── frontend/               # Next.js frontend
│   ├── app/                # Next.js App Router
│   ├── components/         # React components
│   ├── lib/                # Utilities and API client
│   └── types/              # TypeScript types
├── helm/                   # Helm chart
├── docs/                   # Documentation
└── Makefile               # Build automation
```

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/gravitek/crossplane-spy.git
cd crossplane-spy
```

### 2. Backend Development

```bash
# Install dependencies
cd backend
go mod download

# Run backend server
go run cmd/server/main.go

# Or use make
make run-backend
```

The backend API will be available at `http://localhost:8080`.

### 3. Frontend Development

```bash
# Install dependencies
cd frontend
npm install

# Create environment file
cp .env.local.example .env.local

# Run development server
npm run dev

# Or use make
make run-frontend
```

The frontend will be available at `http://localhost:3000`.

### 4. Run Both Services

```bash
# From project root
make dev
```

## Development Workflow

### Adding New API Endpoints

1. Define handler in `backend/internal/api/handlers.go`
2. Register route in `backend/internal/api/router.go`
3. Add client method in `frontend/lib/api.ts`
4. Create TypeScript types in `frontend/types/crossplane.ts`

### Adding shadcn/ui Components

```bash
cd frontend
npx shadcn@latest add <component-name>
```

Example:
```bash
npx shadcn@latest add button
npx shadcn@latest add table
npx shadcn@latest add badge
```

### Code Style

#### Backend (Go)
- Follow standard Go conventions
- Use `gofmt` for formatting
- Run `go vet` before committing
- Document exported functions

```bash
make lint-backend
```

#### Frontend (TypeScript/React)
- Use ESLint configuration
- Follow React best practices
- Use TypeScript strict mode
- Document complex components

```bash
make lint-frontend
```

## Testing

### Backend Tests

```bash
cd backend
go test -v ./...

# With coverage
go test -v -cover ./...
```

### Frontend Tests

```bash
cd frontend
npm test
```

## Building

### Local Build

```bash
# Build everything
make build

# Build backend only
make build-backend

# Build frontend only
make build-frontend
```

### Docker Build

```bash
# Build Docker image
make docker-build

# Run in Docker
make docker-run
```

## Debugging

### Backend Debugging

Use your IDE's Go debugger or add logging:

```go
import "log"

log.Printf("Debug info: %+v", variable)
```

### Frontend Debugging

Use React DevTools and browser console:

```typescript
console.log('Debug info:', variable);
```

## Kubernetes Development

### Testing Against Local Cluster

```bash
# Using kind
kind create cluster --name crossplane-test

# Install Crossplane
helm install crossplane crossplane-stable/crossplane --namespace crossplane-system --create-namespace

# Run crossplane-spy
make dev
```

### Viewing Logs

```bash
# Backend logs
kubectl logs -f deployment/crossplane-spy -c backend

# Frontend logs
kubectl logs -f deployment/crossplane-spy -c frontend
```

## Common Tasks

```bash
# View all available commands
make help

# Install dependencies
make deps

# Run linters
make lint

# Clean build artifacts
make clean

# Run tests
make test
```

## Troubleshooting

### Backend can't connect to Kubernetes

- Ensure `KUBECONFIG` is set correctly
- Check cluster access: `kubectl cluster-info`
- Verify RBAC permissions

### Frontend can't reach backend

- Check `NEXT_PUBLIC_API_URL` in `.env.local`
- Ensure backend is running on correct port
- Check CORS configuration

### Dependencies issues

```bash
# Reset backend dependencies
cd backend
rm go.sum
go mod tidy

# Reset frontend dependencies
cd frontend
rm -rf node_modules package-lock.json
npm install
```
