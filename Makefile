.PHONY: help dev build install clean test test-coverage lint fmt vet check release release-dry release-snapshot version deps tidy verify commit changelog

# Variables
BINARY_NAME=docktidy
MAIN_PATH=./cmd/docktidy
BUILD_DIR=./tmp
DIST_DIR=./dist

# Version info
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Build flags
LDFLAGS=-ldflags "-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"

# Colors for output
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[0;33m
BLUE=\033[0;34m
NC=\033[0m # No Color

## help: Display this help message
help:
	@echo "$(BLUE)docktidy - Development Makefile$(NC)"
	@echo ""
	@echo "$(GREEN)Available targets:$(NC)"
	@grep -E '^## ' $(MAKEFILE_LIST) | sed -e 's/## //' | awk 'BEGIN {FS = ":"}; {printf "  $(YELLOW)%-20s$(NC) %s\n", $$1, $$2}'

## dev: Run with hot reload using Air
dev:
	@echo "$(GREEN)Starting development server with Air...$(NC)"
	@which air > /dev/null || (echo "$(RED)Air not found. Installing...$(NC)" && go install github.com/air-verse/air@latest)
	@air

## build: Build the binary
build:
	@echo "$(GREEN)Building $(BINARY_NAME)...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "$(GREEN)✓ Binary built: $(BUILD_DIR)/$(BINARY_NAME)$(NC)"

## install: Build and install to GOPATH/bin
install:
	@echo "$(GREEN)Installing $(BINARY_NAME)...$(NC)"
	@go install $(LDFLAGS) $(MAIN_PATH)
	@echo "$(GREEN)✓ Installed to $(shell go env GOPATH)/bin/$(BINARY_NAME)$(NC)"

## clean: Clean build artifacts
clean:
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	@rm -rf $(BUILD_DIR) $(DIST_DIR)
	@go clean
	@echo "$(GREEN)✓ Cleaned$(NC)"

## test: Run tests
test:
	@echo "$(GREEN)Running tests...$(NC)"
	@go test -v -race ./...

## test-coverage: Run tests with coverage report
test-coverage:
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	@go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)✓ Coverage report generated: coverage.html$(NC)"

## lint: Run linter
lint:
	@echo "$(GREEN)Running linter...$(NC)"
	@which golangci-lint > /dev/null || (echo "$(RED)golangci-lint not found. Install from https://golangci-lint.run/$(NC)" && exit 1)
	@golangci-lint run ./...

## fmt: Format code
fmt:
	@echo "$(GREEN)Formatting code...$(NC)"
	@go fmt ./...
	@echo "$(GREEN)✓ Code formatted$(NC)"

## vet: Run go vet
vet:
	@echo "$(GREEN)Running go vet...$(NC)"
	@go vet ./...
	@echo "$(GREEN)✓ No issues found$(NC)"

## check: Run all checks (fmt, vet, lint, test)
check: fmt vet lint test
	@echo "$(GREEN)✓ All checks passed!$(NC)"

## tidy: Tidy dependencies
tidy:
	@echo "$(GREEN)Tidying dependencies...$(NC)"
	@go mod tidy
	@echo "$(GREEN)✓ Dependencies tidied$(NC)"

## deps: Download dependencies
deps:
	@echo "$(GREEN)Downloading dependencies...$(NC)"
	@go mod download
	@echo "$(GREEN)✓ Dependencies downloaded$(NC)"

## verify: Verify dependencies
verify:
	@echo "$(GREEN)Verifying dependencies...$(NC)"
	@go mod verify
	@echo "$(GREEN)✓ Dependencies verified$(NC)"

## commit: Create a conventional commit interactively
commit:
	@which cog > /dev/null || (echo "$(RED)Cocogitto not found. Install with: brew install cocogitto$(NC)" && exit 1)
	@cog commit

## changelog: Generate changelog
changelog:
	@echo "$(GREEN)Generating changelog...$(NC)"
	@which cog > /dev/null || (echo "$(RED)Cocogitto not found. Install with: brew install cocogitto$(NC)" && exit 1)
	@cog changelog
	@echo "$(GREEN)✓ Changelog generated$(NC)"

## release: Create a new release (auto-detect version)
release:
	@echo "$(GREEN)Creating release...$(NC)"
	@which cog > /dev/null || (echo "$(RED)Cocogitto not found. Install with: brew install cocogitto$(NC)" && exit 1)
	@echo "$(YELLOW)Checking commits...$(NC)"
	@cog check || (echo "$(RED)✗ Commit check failed. Fix commits before releasing.$(NC)" && exit 1)
	@echo "$(YELLOW)Bumping version...$(NC)"
	@cog bump --auto
	@echo "$(GREEN)✓ Release created and pushed!$(NC)"

## release-dry: Dry run release (show what would happen)
release-dry:
	@echo "$(GREEN)Dry run release...$(NC)"
	@which cog > /dev/null || (echo "$(RED)Cocogitto not found. Install with: brew install cocogitto$(NC)" && exit 1)
	@cog bump --auto --dry-run

## release-snapshot: Build release locally without publishing
release-snapshot:
	@echo "$(GREEN)Building release snapshot...$(NC)"
	@which goreleaser > /dev/null || (echo "$(RED)GoReleaser not found. Install with: brew install goreleaser$(NC)" && exit 1)
	@goreleaser release --snapshot --clean
	@echo "$(GREEN)✓ Release snapshot built in $(DIST_DIR)/$(NC)"

## version: Show version information
version:
	@echo "$(BLUE)Version:$(NC)  $(VERSION)"
	@echo "$(BLUE)Commit:$(NC)   $(COMMIT)"
	@echo "$(BLUE)Built:$(NC)    $(DATE)"

## run: Build and run the binary
run: build
	@echo "$(GREEN)Running $(BINARY_NAME)...$(NC)"
	@$(BUILD_DIR)/$(BINARY_NAME)

## docker-check: Check Docker connection
docker-check:
	@echo "$(GREEN)Checking Docker connection...$(NC)"
	@docker info > /dev/null 2>&1 && echo "$(GREEN)✓ Docker is running$(NC)" || (echo "$(RED)✗ Docker is not running$(NC)" && exit 1)

## tools: Install development tools
tools:
	@echo "$(GREEN)Installing development tools...$(NC)"
	@echo "$(YELLOW)Installing Air...$(NC)"
	@go install github.com/air-verse/air@latest
	@echo "$(YELLOW)Installing golangci-lint...$(NC)"
	@brew install golangci-lint 2>/dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin
	@echo "$(YELLOW)Installing Cocogitto...$(NC)"
	@brew install cocogitto 2>/dev/null || cargo install cocogitto
	@echo "$(YELLOW)Installing GoReleaser...$(NC)"
	@brew install goreleaser 2>/dev/null || go install github.com/goreleaser/goreleaser@latest
	@echo "$(GREEN)✓ All tools installed$(NC)"

## setup: Initial project setup (install tools and dependencies)
setup: tools deps
	@echo "$(GREEN)Setting up git hooks...$(NC)"
	@cog install-hook --all 2>/dev/null || (echo "$(YELLOW)⚠ Could not install hooks automatically. Run 'cog install-hook --all' manually.$(NC)")
	@echo "$(GREEN)✓ Project setup complete!$(NC)"
	@echo ""
	@echo "$(BLUE)Quick start:$(NC)"
	@echo "  make dev      - Start development server"
	@echo "  make test     - Run tests"
	@echo "  make build    - Build binary"

## ci: Run CI checks locally (mimics GitHub Actions)
ci: tidy verify check
	@echo "$(GREEN)✓ CI checks passed!$(NC)"

# Default target
.DEFAULT_GOAL := help
