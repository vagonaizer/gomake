package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

// MakefileGenerator handles Makefile generation
type MakefileGenerator struct {
	config *Config
	logger Logger
}

// NewMakefileGenerator creates a new Makefile generator
func NewMakefileGenerator(config *Config, logger Logger) *MakefileGenerator {
	return &MakefileGenerator{
		config: config,
		logger: logger,
	}
}

// Generate creates a Makefile
func (mg *MakefileGenerator) Generate(projectPath string) error {
	content := fmt.Sprintf(`# %s Makefile

.PHONY: help build run test clean lint fmt install

APP_NAME=%s
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%%-15s\033[0m %%s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go
	@echo "Build complete: bin/$(APP_NAME)"

run: ## Run the application
	@go run cmd/$(APP_NAME)/main.go

test: ## Run tests
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

clean: ## Clean build artifacts
	@rm -rf bin/
	@rm -f coverage.out coverage.html

lint: ## Run linter
	@golangci-lint run

fmt: ## Format code
	@go fmt ./...
	@goimports -w .

mod-tidy: ## Tidy go modules
	@go mod tidy

deps: ## Download dependencies
	@go mod download

docker-build: ## Build Docker image
	@docker build -t $(APP_NAME):$(VERSION) .

docker-run: ## Run Docker container
	@docker run -p 8080:8080 $(APP_NAME):$(VERSION)

dev: ## Run in development mode with hot reload
	@air

install-tools: ## Install development tools
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/cosmtrek/air@latest

setup: install-tools deps ## Setup development environment
	@echo "Development environment setup complete"

all: fmt lint test build ## Run all checks and build

.DEFAULT_GOAL := help
`, mg.config.ProjectName, mg.config.ProjectName)

	filePath := filepath.Join(projectPath, "Makefile")
	return os.WriteFile(filePath, []byte(content), 0644)
}
