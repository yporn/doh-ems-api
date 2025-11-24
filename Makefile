# Makefile
.PHONY: help dev dev-simple test watch-test build clean install-deps

# Default target
help:
	@echo "Available commands:"
	@echo "  make dev          - Start development server with hot reload (full config)"
	@echo "  make dev-simple   - Start development server (simple mode)"
	@echo "  make test         - Run tests once"
	@echo "  make watch-test   - Run tests with hot reload"
	@echo "  make build        - Build the application"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make install-deps - Install development dependencies"

# Development with full reflex config
dev:
	@echo "🔥 Starting development server with hot reload (single watcher)..."
	@reflex -c reflex.conf

# Simple development mode
dev-simple:
	@echo "🔥 Starting simple hot reload..."
	@reflex -r '\.go$$' -s --decoration=fancy -- go run cmd/app/main.go

# Run tests once
test:
	@echo "🧪 Running tests..."
	@go test -v ./...

# Watch tests
watch-test:
	@echo "🧪 Watching tests..."
	@reflex -r '_test\.go$$' --decoration=fancy -- go test -v ./...

# Build application
build:
	@echo "🔨 Building application..."
	@mkdir -p bin
	@go build -o bin/api cmd/app/main.go
	@echo "✅ Build complete: bin/app"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	@rm -rf bin/ tmp/
	@echo "✅ Clean complete"

# Install development dependencies
install-deps:
	@echo "📦 Installing development dependencies..."
	@go install github.com/cespare/reflex@latest
	@echo "✅ Dependencies installed"

# Quick run without hot reload
run:
	@mkdir -p bin
	@go build -trimpath -o bin/doh-ems ./cmd/app && ./bin/doh-ems

# Format code
fmt:
	@echo "🎨 Formatting code..."
	@go fmt ./...
	@echo "✅ Code formatted"

# Run linter
lint:
	@echo "🔍 Running linter..."
	@golangci-lint run
	@echo "✅ Linting complete"

# Watch and lint
watch-lint:
	@echo "🔍 Watching and linting..."
	@reflex -r '\.go$$' -- golangci-lint run --fast