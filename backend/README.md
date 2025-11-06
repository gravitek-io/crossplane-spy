# Crossplane Spy - Backend

Go backend service for Crossplane Spy dashboard.

## Structure

- `cmd/server/` - Main application entry point
- `internal/api/` - HTTP API handlers and routing
- `internal/k8s/` - Kubernetes client and resource discovery
- `internal/models/` - Data models for Crossplane resources
- `pkg/` - Public packages (if needed)

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [client-go](https://github.com/kubernetes/client-go) - Kubernetes client library

## Development

### Prerequisites

- Go 1.23 or later
- Access to a Kubernetes cluster with Crossplane installed

### Running locally

```bash
# Install dependencies
go mod download

# Run the server
go run cmd/server/main.go
```

The API will be available at `http://localhost:8080`.

### Building

```bash
go build -o crossplane-spy cmd/server/main.go
```

## API Endpoints

- `GET /health` - Health check
- `GET /api/v1/providers` - List all Providers
- `GET /api/v1/providerconfigs` - List all ProviderConfigs
- `GET /api/v1/xrds` - List all XRDs
- `GET /api/v1/compositions` - List all Compositions
- `GET /api/v1/xrs` - List all Composite Resources
- `GET /api/v1/functions` - List all Functions
- `GET /api/v1/cluster-resources` - List cluster-scoped resources
- `GET /api/v1/namespace-resources` - List namespace-scoped resources

## Configuration

- `PORT` - Server port (default: 8080)
- `KUBECONFIG` - Path to kubeconfig file (default: ~/.kube/config)
