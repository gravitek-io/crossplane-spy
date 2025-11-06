# Crossplane Spy - Project Guide for AI Assistants

This document provides context and guidelines for AI assistants working on the Crossplane Spy project.

## Project Overview

**Crossplane Spy** is a web dashboard for visualizing Crossplane v2 resources in Kubernetes clusters. It's designed as an educational tool for platform teams learning Crossplane, with a focus on understanding cluster-scoped vs namespace-scoped resources.

### Key Goals
- **Educational**: Help users understand Crossplane v2 concepts and multi-tenancy
- **Simple**: KISS principle - keep it simple and maintainable
- **Read-only**: Monitor and visualize, don't modify resources
- **Internal tool**: For platform teams, not end-users

## Architecture

### Tech Stack

**Backend (Go 1.23+)**
- Framework: Gin
- Kubernetes client: client-go
- Purpose: REST API + serves frontend static files (single binary deployment)

**Frontend (Next.js 15)**
- Framework: Next.js with App Router
- Styling: TailwindCSS + shadcn/ui
- Language: TypeScript (strict mode)
- Build: Standalone output for Docker

**Deployment**
- Single Docker image (multi-stage build)
- Helm chart for Kubernetes deployment
- Namespace: `crossplane-spy` (dedicated)
- RBAC: Read-only ClusterRole

### Project Structure

```
crossplane-spy/
├── backend/
│   ├── cmd/server/        # Main application entry point
│   ├── internal/
│   │   ├── api/          # REST API handlers (Gin)
│   │   ├── k8s/          # Kubernetes client & resource discovery
│   │   └── models/       # Data models & converters
│   └── pkg/              # Public packages (if needed)
├── frontend/
│   ├── app/              # Next.js App Router pages
│   ├── components/       # React components
│   │   ├── nav/         # Navigation (sidebar, header)
│   │   ├── resources/   # Resource listing components
│   │   └── ui/          # UI primitives (badges, etc.)
│   ├── lib/             # Utilities (API client, date utils)
│   └── types/           # TypeScript type definitions
├── helm/crossplane-spy/  # Helm chart for deployment
├── docs/                # Documentation
└── Makefile            # Development commands
```

## Development Guidelines

### Code Style

**Go**
- Follow standard Go conventions
- Use `gofmt` and `go vet`
- Document exported functions and types
- Keep packages focused and cohesive
- Error handling: always wrap errors with context

**TypeScript/React**
- Use functional components with hooks
- Strict TypeScript mode enabled
- Follow ESLint configuration
- Use shadcn/ui components for consistency
- Server components by default, "use client" only when needed

### Important Patterns

**Backend - Resource Discovery**
- Use dynamic client for custom resources (XRs)
- Discover GVRs from XRDs at runtime
- Detect scope (cluster/namespace) automatically
- Handle missing resources gracefully (return empty lists)

**Frontend - Data Fetching**
- Client-side fetching with `useEffect` in "use client" components
- Show loading states
- Handle errors with user-friendly messages
- Use the `ResourceListPage` component for consistency

### Common Issues & Solutions

**TypeScript errors with lucide-react**
- Solution: Restart TS server in VSCode (`Cmd+Shift+P` → "TypeScript: Restart TS Server")
- Root cause: VSCode caches types

**Backend can't connect to cluster**
- Check: `~/.kube/config` is configured
- The app uses the same kubeconfig as `kubectl`
- For in-cluster deployment, it auto-detects and uses ServiceAccount

**Frontend build errors**
- Ensure `autoprefixer` is in devDependencies
- Run `npm install` in `frontend/` directory

## Development Workflow

### Local Development

```bash
# Terminal 1 - Backend
cd backend
go run cmd/server/main.go
# Runs on http://localhost:8080

# Terminal 2 - Frontend
cd frontend
npm run dev
# Runs on http://localhost:3000

# Or use Makefile
make dev  # Runs both in parallel
```

### Testing

```bash
# Backend tests
cd backend && go test -v ./...

# Frontend linting
cd frontend && npm run lint

# Type checking
cd frontend && npx tsc --noEmit
```

### Building

```bash
# Build both
make build

# Docker image (single image with both backend + frontend)
make docker-build

# Helm install
helm install crossplane-spy ./helm/crossplane-spy
```

## API Endpoints

Backend serves both API and frontend:

**API (JSON)**
- `GET /health` - Health check
- `GET /api/v1/resources` - Summary of all resources
- `GET /api/v1/providers` - List Providers
- `GET /api/v1/xrds` - List XRDs
- `GET /api/v1/compositions` - List Compositions
- `GET /api/v1/functions` - List Functions
- `GET /api/v1/providerconfigs` - List ProviderConfigs
- `GET /api/v1/xrs` - List Composite Resources
- `GET /api/v1/cluster-resources` - All cluster-scoped resources
- `GET /api/v1/namespace-resources` - All namespace-scoped resources

**Frontend (HTML)**
- `/*` - Serves Next.js static files (SPA)
- `/_next/*` - Next.js assets

## Key Files to Understand

**Backend**
- `backend/internal/k8s/client.go` - Kubernetes client setup (in-cluster vs kubeconfig)
- `backend/internal/k8s/resources.go` - Resource listing methods
- `backend/internal/k8s/discovery.go` - GVR discovery for XRs and ProviderConfigs
- `backend/internal/models/converter.go` - Convert K8s unstructured to typed models
- `backend/internal/api/handlers.go` - API endpoint handlers

**Frontend**
- `frontend/app/layout.tsx` - Root layout with sidebar navigation
- `frontend/components/resources/resource-list-page.tsx` - Reusable page for listing resources
- `frontend/components/resources/resource-table.tsx` - Table component with scope/status badges
- `frontend/lib/api.ts` - API client for backend communication
- `frontend/types/crossplane.ts` - TypeScript types for Crossplane resources

## Important Constraints

### Security
- **Read-only access**: Never modify Kubernetes resources
- **RBAC**: ClusterRole grants only `get`, `list`, `watch` verbs
- **No authentication**: Designed for internal cluster use (not exposed externally by default)
- **Non-root**: Docker image runs as UID 1000

### Scope Handling
- Crossplane v2 introduced namespace-scoped resources
- The app must clearly distinguish between cluster and namespace scope
- This is a key educational feature for understanding multi-tenancy

### Performance
- Resource discovery can be slow with many XRDs
- Don't block on slow queries - return partial results
- Consider caching for future enhancements (not implemented yet)

## Contributing Guidelines

### Before Making Changes

1. **Understand the context**: This is an educational tool, simplicity > features
2. **Check existing patterns**: Follow established code patterns
3. **Test locally**: Run both backend and frontend before committing
4. **Documentation**: Update docs if adding features

### Commit Message Format

Follow Conventional Commits:
- `feat:` - New feature
- `fix:` - Bug fix
- `docs:` - Documentation changes
- `refactor:` - Code refactoring
- `test:` - Adding tests
- `chore:` - Maintenance tasks

Example: `feat: add filtering by namespace in resource tables`

**Important**: No Claude/AI references in commit messages (per user's global config)

### When Adding New Resource Types

1. Add GVR constant in `backend/internal/k8s/resources.go`
2. Add listing method in `backend/internal/k8s/resources.go`
3. Add model in `backend/internal/models/resources.go`
4. Add API endpoint in `backend/internal/api/router.go` and handler in `handlers.go`
5. Add API client method in `frontend/lib/api.ts`
6. Add TypeScript type in `frontend/types/crossplane.ts`
7. Add page in `frontend/app/<resource-type>/page.tsx`
8. Add navigation link in `frontend/components/nav/sidebar.tsx`

### When Adding UI Components

- Use shadcn/ui primitives: `npx shadcn@latest add <component>`
- Follow the `components/ui/` structure for reusable components
- Keep components small and focused
- Document props with TypeScript interfaces
- Use TailwindCSS for styling (no CSS modules)

## Debugging Tips

### Backend Issues

```bash
# Check logs
go run cmd/server/main.go

# Test API directly
curl http://localhost:8080/health
curl http://localhost:8080/api/v1/resources

# Check cluster access
kubectl get providers
kubectl get xrds
```

### Frontend Issues

```bash
# Check build
npm run build

# Type check
npx tsc --noEmit

# Inspect API calls in browser DevTools Network tab
```

### Common Problems

**"Failed to fetch"** in frontend
- Backend not running → Start it with `make run-backend`
- Wrong API URL → Check `.env.local` has correct `NEXT_PUBLIC_API_URL`

**Empty resource lists**
- No Crossplane in cluster → Expected, app works but shows empty state
- RBAC issues → Check ClusterRole bindings

**Build failures**
- Missing dependencies → Run `npm install` or `go mod tidy`
- Type errors → Restart TypeScript server in IDE

## Future Enhancements (Not Implemented)

Ideas for future development:
- WebSocket/SSE for real-time updates
- Resource relationship graph visualization (cytoscape.js)
- Filtering by labels, annotations
- Resource detail pages with YAML view
- Export functionality (JSON, YAML)
- Dark mode toggle
- Multi-cluster support

## License & Credits

- License: Apache 2.0
- No Claude references in code or commits
- Clean, professional code for open-source community

## Questions?

Refer to:
- [README.md](README.md) - User documentation
- [docs/architecture.md](docs/architecture.md) - Architecture details
- [docs/development.md](docs/development.md) - Development guide
- [CONTRIBUTING.md](CONTRIBUTING.md) - Contribution guidelines
