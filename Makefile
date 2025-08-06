# gomake Makefile - Simplified version

.PHONY: help build run test clean lint fmt install

# Variables
APP_NAME=gomake
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X github.com/gomake/internal/generator.Version=$(VERSION)"

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@go build $(LDFLAGS) -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go
	@echo "Build complete: bin/$(APP_NAME)"

run: ## Run the application
	@go run cmd/$(APP_NAME)/main.go

test: ## Run tests
	@go test -v ./...

clean: ## Clean build artifacts
	@rm -rf bin/

lint: ## Run linter
	@golangci-lint run

fmt: ## Format code
	@go fmt ./...
	@goimports -w .

install: build ## Install the binary
	@cp bin/$(APP_NAME) /usr/local/bin/
	@echo "$(APP_NAME) installed to /usr/local/bin/"

mod-tidy: ## Tidy go modules
	@go mod tidy

all: fmt lint test build ## Run all checks and build

.DEFAULT_GOAL := help
