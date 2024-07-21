.PHONY: all build run swag docker-up docker-down clean test

# Define variables
APP_NAME=project-management
DOCKER_COMPOSE_FILE=docker-compose.yml
DOCKER_SERVICE=app

# Default target
all: build

# Build the Go application
build:
	go build -o ./tmp/main ./cmd/main.go

# Run the Go application
run: build
	./tmp/main

# Generate Swagger documentation
swag:
	swag init -g cmd/main.go

# Build and run Docker containers
docker-up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up --build

# Stop Docker containers
docker-down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Test Unit tests in docker containers
docker-test:
	docker compose -f ${DOCKER_COMPOSE_FILE} run --rm app make test

# Clean up build artifacts
clean:
	rm -rf tmp

# Run unit tests
test:
	go test ./internal/... -v

# Build and run the application with Air (includes Swagger generation)
air:
	./generate_docs.sh && air -c .air.toml
