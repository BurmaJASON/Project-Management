#!/bin/sh

# Run swag init to generate Swagger documentation
swag init -g cmd/main.go

# Build the Go application
go build -o ./tmp/main ./cmd/main.go
