.PHONY: build run test lint clean migrate-up migrate-down

# Build variables
BINARY_NAME=app
BUILD_DIR=./bin

# Go variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

# Database variables
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=app
DB_SSL_MODE=disable
MIGRATION_DIR=./scripts/migrations

# Build the application
build:
	@echo "Building application..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/app

# Run the application
run:
	@echo "Running application..."
	@go run ./cmd/app

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

# Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Run database migrations up
migrate-up:
	@echo "Running migrations up..."
	@migrate -path $(MIGRATION_DIR) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" up

# Run database migrations down
migrate-down:
	@echo "Running migrations down..."
	@migrate -path $(MIGRATION_DIR) -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" down

# Create a new migration file
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $$name

# Generate API documentation
docs:
	@echo "Generating API documentation..."
	@swag init -g cmd/app/main.go -o api/docs

# Run the application in development mode
dev:
	@echo "Running in development mode..."
	@go run ./cmd/app

# Help command
help:
	@echo "Available commands:"
	@echo "  make build          - Build the application"
	@echo "  make run            - Run the application"
	@echo "  make test           - Run tests"
	@echo "  make test-coverage  - Run tests with coverage"
	@echo "  make lint           - Run linter"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make deps           - Install dependencies"
	@echo "  make migrate-up     - Run database migrations up"
	@echo "  make migrate-down   - Run database migrations down"
	@echo "  make migrate-create - Create a new migration file"
	@echo "  make docs           - Generate API documentation"
	@echo "  make dev            - Run the application in development mode"
	@echo "  make help           - Show this help message" 