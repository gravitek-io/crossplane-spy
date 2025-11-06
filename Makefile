# Crossplane Spy - Makefile
# Simplifies common development and deployment tasks

.PHONY: help dev build test clean docker-build docker-push install-frontend run-backend run-frontend

# Variables
BACKEND_BINARY := backend/crossplane-spy
DOCKER_IMAGE := crossplane-spy
DOCKER_TAG := latest
DOCKER_REGISTRY ?= # Set your registry here

help: ## Show this help message
	@echo "Crossplane Spy - Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Development commands
dev: ## Start both backend and frontend in development mode
	@echo "Starting backend and frontend..."
	@make -j2 run-backend run-frontend

install-frontend: ## Install frontend dependencies
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

run-backend: ## Run backend in development mode
	@echo "Starting backend server..."
	cd backend && go run cmd/server/main.go

run-frontend: install-frontend ## Run frontend in development mode
	@echo "Starting frontend development server..."
	cd frontend && npm run dev

# Build commands
build: build-backend build-frontend ## Build both backend and frontend

build-backend: ## Build backend binary
	@echo "Building backend..."
	cd backend && go build -o crossplane-spy cmd/server/main.go

build-frontend: install-frontend ## Build frontend for production
	@echo "Building frontend..."
	cd frontend && npm run build

# Testing
test: test-backend ## Run all tests
	@echo "All tests completed"

test-backend: ## Run backend tests
	@echo "Running backend tests..."
	cd backend && go test -v ./...

# Docker commands
docker-build: ## Build Docker image
	@echo "Building Docker image $(DOCKER_IMAGE):$(DOCKER_TAG)..."
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

docker-push: docker-build ## Build and push Docker image
	@echo "Pushing Docker image to $(DOCKER_REGISTRY)$(DOCKER_IMAGE):$(DOCKER_TAG)..."
	docker tag $(DOCKER_IMAGE):$(DOCKER_TAG) $(DOCKER_REGISTRY)$(DOCKER_IMAGE):$(DOCKER_TAG)
	docker push $(DOCKER_REGISTRY)$(DOCKER_IMAGE):$(DOCKER_TAG)

docker-run: docker-build ## Run application in Docker container
	@echo "Running Docker container..."
	docker run -p 8080:8080 -p 3000:3000 $(DOCKER_IMAGE):$(DOCKER_TAG)

# Helm commands
helm-install: ## Install application using Helm
	@echo "Installing with Helm..."
	helm install crossplane-spy ./helm/crossplane-spy

helm-upgrade: ## Upgrade application using Helm
	@echo "Upgrading with Helm..."
	helm upgrade crossplane-spy ./helm/crossplane-spy

helm-uninstall: ## Uninstall application using Helm
	@echo "Uninstalling with Helm..."
	helm uninstall crossplane-spy

# Cleanup
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	rm -f $(BACKEND_BINARY)
	rm -rf frontend/.next
	rm -rf frontend/out
	rm -rf frontend/node_modules

# Linting and formatting
lint: lint-backend lint-frontend ## Run all linters

lint-backend: ## Run backend linter
	@echo "Running backend linter..."
	cd backend && go vet ./...
	cd backend && go fmt ./...

lint-frontend: ## Run frontend linter
	@echo "Running frontend linter..."
	cd frontend && npm run lint

# Dependencies
deps: deps-backend install-frontend ## Update all dependencies

deps-backend: ## Update backend dependencies
	@echo "Updating backend dependencies..."
	cd backend && go mod tidy
	cd backend && go mod download
