.PHONY: all build test test-cover lint fmt vet clean help

# Variables
BINARY_NAME=gtvapi
GO=go
GOFLAGS=
GOFILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Default target
all: fmt vet lint test build

# Build the project
build:
	@echo "Building..."
	$(GO) build $(GOFLAGS) ./...

# Run tests
test:
	@echo "Running tests..."
	$(GO) test -v -race ./...

# Run tests with coverage
test-cover:
	@echo "Running tests with coverage..."
	$(GO) test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run tests with coverage and display percentage
cover: test-cover
	@echo "Coverage summary:"
	@$(GO) tool cover -func=coverage.out | grep total | awk '{print "Total coverage: " $$3}'

# Run linter
lint:
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Install with: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin"; \
	fi

# Format code
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	$(GO) vet ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GO) clean
	rm -f coverage.out coverage.html

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	$(GO) test -bench=. -benchmem ./...

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GO) mod download
	$(GO) mod tidy

# Update dependencies
update-deps:
	@echo "Updating dependencies..."
	$(GO) get -u ./...
	$(GO) mod tidy

# Run example
example:
	@echo "Running example..."
	@$(GO) run examples/*.go

# Generate documentation
doc:
	@echo "Opening documentation in browser..."
	@echo "Visit: https://pkg.go.dev/github.com/wehmoen-dev/gtvapi"

# Help target
help:
	@echo "Available targets:"
	@echo "  all         - Format, vet, lint, test, and build (default)"
	@echo "  build       - Build the project"
	@echo "  test        - Run tests"
	@echo "  test-cover  - Run tests with coverage report"
	@echo "  cover       - Run tests with coverage and display percentage"
	@echo "  lint        - Run golangci-lint"
	@echo "  fmt         - Format code with go fmt"
	@echo "  vet         - Run go vet"
	@echo "  bench       - Run benchmarks"
	@echo "  deps        - Download dependencies"
	@echo "  update-deps - Update dependencies"
	@echo "  clean       - Clean build artifacts"
	@echo "  doc         - Show documentation URL"
	@echo "  help        - Show this help message"
