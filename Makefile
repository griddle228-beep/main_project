.PHONY: help build run test migrate-up migrate-down migrate-create docker-build docker-up docker-down clean

# Variables
APP_NAME=devx-service-backend
MIGRATIONS_DIR=./migrations
DOCKER_COMPOSE=docker-compose

help: ## Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	go build -o bin/$(APP_NAME) main.go
	go build -o bin/migrator cmd/migrator/main.go

run: ## Run the application locally
	@echo "Running $(APP_NAME)..."
	go run cmd/app/main.go

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

deps: ## Install dependencies
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

migrate-up: ## Run database migrations up
	@echo "Running migrations up..."
	go run cmd/migrator/main.go $(MIGRATIONS_DIR) up

migrate-down: ## Rollback last migration
	@echo "Rolling back last migration..."
	go run cmd/migrator/main.go $(MIGRATIONS_DIR) down

migrate-create: ## Create a new migration file (usage: make migrate-create NAME=add_users_table)
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make migrate-create NAME=your_migration_name"; \
		exit 1; \
	fi
	@NEXT_NUM=$$(ls -1 $(MIGRATIONS_DIR)/*.sql 2>/dev/null | wc -l | xargs); \
	NEXT_NUM=$$((NEXT_NUM + 1)); \
	FILE_NUM=$$(printf "%05d" $$NEXT_NUM); \
	FILENAME="$(MIGRATIONS_DIR)/$${FILE_NUM}_$(NAME).sql"; \
	echo "Creating migration: $$FILENAME"; \
	echo "-- +goose Up" > $$FILENAME; \
	echo "-- +goose StatementBegin" >> $$FILENAME; \
	echo "-- Write your UP migration here" >> $$FILENAME; \
	echo "-- +goose StatementEnd" >> $$FILENAME; \
	echo "" >> $$FILENAME; \
	echo "-- +goose Down" >> $$FILENAME; \
	echo "-- +goose StatementBegin" >> $$FILENAME; \
	echo "-- Write your DOWN migration here" >> $$FILENAME; \
	echo "-- +goose StatementEnd" >> $$FILENAME; \
	echo "Migration created: $$FILENAME"

docker-build: ## Build Docker images
	@echo "Building Docker images..."
	$(DOCKER_COMPOSE) build

docker-up: ## Start Docker containers with migrations
	@echo "Starting Docker containers..."
	$(DOCKER_COMPOSE) up -d
	@echo "Waiting for database to be ready..."
	@sleep 5
	@echo "Running migrations..."
	@$(DOCKER_COMPOSE) exec -T backend go run cmd/migrator/main.go ./migrations up || true
	@echo "Application is ready!"

docker-down: ## Stop Docker containers
	@echo "Stopping Docker containers..."
	$(DOCKER_COMPOSE) down

docker-logs: ## Show Docker logs
	$(DOCKER_COMPOSE) logs -f

docker-restart: docker-down docker-up ## Restart Docker containers

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/
	go clean

lint: ## Run linter
	@echo "Running linter..."
	golangci-lint run

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...
	gofmt -s -w .

dev: ## Run in development mode with auto-reload (requires air)
	@echo "Running in development mode..."
	air

install-tools: ## Install development tools
	@echo "Installing development tools..."
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.DEFAULT_GOAL := help
