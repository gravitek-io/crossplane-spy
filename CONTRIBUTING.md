# Contributing to Crossplane Spy

Thank you for considering contributing to Crossplane Spy! This document provides guidelines and instructions for contributing.

## Code of Conduct

This project is an educational tool for the Crossplane community. Please be respectful and constructive in all interactions.

## How to Contribute

### Reporting Bugs

If you find a bug, please open an issue with:
- A clear title and description
- Steps to reproduce the issue
- Expected vs actual behavior
- Your environment (Kubernetes version, Crossplane version, etc.)

### Suggesting Features

Feature requests are welcome! Please:
- Check if the feature has already been requested
- Provide a clear use case
- Explain how it helps understanding Crossplane concepts

### Pull Requests

1. **Fork the repository**

2. **Create a feature branch**
   ```bash
   git checkout -b feature/amazing-feature
   ```

3. **Make your changes**
   - Follow the existing code style
   - Add documentation for new features
   - Ensure code is well-commented
   - Write tests when applicable

4. **Test your changes**
   ```bash
   # Backend tests
   cd backend && go test ./...

   # Frontend lint
   cd frontend && npm run lint
   ```

5. **Commit your changes**
   ```bash
   git commit -m 'feat: add amazing feature'
   ```

   Use [Conventional Commits](https://www.conventionalcommits.org/):
   - `feat:` new feature
   - `fix:` bug fix
   - `docs:` documentation changes
   - `refactor:` code refactoring
   - `test:` adding tests
   - `chore:` maintenance tasks

6. **Push to your fork**
   ```bash
   git push origin feature/amazing-feature
   ```

7. **Open a Pull Request**
   - Provide a clear description
   - Reference any related issues
   - Explain the changes and their impact

## Development Setup

### Prerequisites

- Go 1.23+
- Node.js 20+
- Docker (for building images)
- Kubernetes cluster with Crossplane (for testing)

### Local Development

```bash
# Install dependencies
make deps

# Run backend
make run-backend

# Run frontend (in another terminal)
make run-frontend

# Run both (if your terminal supports it)
make dev
```

### Testing

```bash
# Backend tests
cd backend && go test -v ./...

# Frontend lint
cd frontend && npm run lint

# Build Docker image
make docker-build
```

## Code Style

### Go

- Follow standard Go conventions
- Use `gofmt` for formatting
- Run `go vet` before committing
- Document exported functions and types

### TypeScript/React

- Use TypeScript strict mode
- Follow ESLint configuration
- Use functional components with hooks
- Document complex components

### Documentation

- Keep documentation up to date
- Use clear, concise language
- Include examples when helpful
- Use mermaid diagrams for architecture

## Project Structure

```
crossplane-spy/
├── backend/           # Go backend
│   ├── cmd/          # Application entry points
│   ├── internal/     # Private application code
│   └── pkg/          # Public packages
├── frontend/         # Next.js frontend
│   ├── app/         # Next.js App Router
│   ├── components/  # React components
│   └── lib/         # Utilities
├── helm/            # Helm chart
└── docs/            # Documentation
```

## Review Process

1. Automated checks must pass (linting, tests)
2. Code review by maintainers
3. Documentation review
4. Testing in a real environment (when applicable)

## Questions?

Feel free to open an issue for questions or reach out to the maintainers.

## License

By contributing, you agree that your contributions will be licensed under the Apache License 2.0.
