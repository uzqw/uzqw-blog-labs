.PHONY: help verify test test-race fmt vet tidy clean build bench

# Find all directories containing go.mod files (blog post folders)
MODULES := $(dir $(shell find . -name "go.mod" -not -path "./vendor/*"))

# Go commands
GOCMD := go
GOTEST := $(GOCMD) test
GOFMT := $(GOCMD) fmt
GOVET := $(GOCMD) vet
GOMOD := $(GOCMD) mod

# Default target
help:
	@echo "uzqw-blog-labs Makefile"
	@echo ""
	@echo "Available targets:"
	@echo "  verify      - Run all CI checks locally (mirrors GitHub Actions)"
	@echo "  test        - Run all tests across blog post projects"
	@echo "  test-race   - Run tests with race detector"
	@echo "  bench       - Run all benchmarks"
	@echo "  fmt         - Format all Go code"
	@echo "  vet         - Run go vet on all projects"
	@echo "  tidy        - Tidy all Go modules"
	@echo "  build       - Build all projects"
	@echo "  clean       - Clean build artifacts and coverage files"
	@echo "  help        - Show this help message"

# Run all verification checks locally (mirrors GitHub Actions CI)
verify: tidy fmt vet build
	@echo "=== Running CI checks locally ==="
	@echo ""
	@for dir in $(MODULES); do \
		echo ">>> Checking $$dir"; \
		(cd $$dir && \
		 echo "  - Running tests with race detector..." && \
		 $(GOTEST) -v -race -coverprofile=coverage.out ./... && \
		 echo "  - Coverage:" && \
		 $(GOCMD) tool cover -func=coverage.out 2>/dev/null | tail -n 1 || echo "  No coverage data" \
		) || exit 1; \
		echo ""; \
	done
	@echo "=== All CI checks passed! ==="

# Run all tests
test:
	@echo "Running tests in all modules..."
	@for dir in $(MODULES); do \
		echo ">>> Testing $$dir"; \
		(cd $$dir && $(GOTEST) -v ./...) || exit 1; \
	done

# Run tests with race detector
test-race:
	@echo "Running tests with race detector..."
	@for dir in $(MODULES); do \
		echo ">>> Testing $$dir"; \
		(cd $$dir && $(GOTEST) -v -race ./...) || exit 1; \
	done

# Run benchmarks
bench:
	@echo "Running benchmarks in all modules..."
	@for dir in $(MODULES); do \
		echo ">>> Benchmarking $$dir"; \
		(cd $$dir && $(GOTEST) -bench=. -benchmem -benchtime=3s ./...) || exit 1; \
	done

# Format code
fmt:
	@echo "Formatting code in all modules..."
	@for dir in $(MODULES); do \
		echo ">>> Formatting $$dir"; \
		(cd $$dir && $(GOFMT) ./...); \
	done

# Run go vet
vet:
	@echo "Running go vet in all modules..."
	@for dir in $(MODULES); do \
		echo ">>> Vetting $$dir"; \
		(cd $$dir && $(GOVET) ./...) || exit 1; \
	done

# Tidy dependencies
tidy:
	@echo "Tidying dependencies in all modules..."
	@for dir in $(MODULES); do \
		echo ">>> Tidying $$dir"; \
		(cd $$dir && $(GOMOD) tidy); \
	done

# Build all projects
build:
	@echo "Building all modules..."
	@for dir in $(MODULES); do \
		echo ">>> Building $$dir"; \
		(cd $$dir && $(GOCMD) build ./...) || exit 1; \
	done

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@find . -name "coverage.out" -delete
	@find . -name "coverage.html" -delete
	@find . -name "*.test" -delete
	@echo "Clean complete!"
