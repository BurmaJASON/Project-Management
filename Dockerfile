FROM golang:1.22-alpine

# Install necessary packages and dependencies
RUN apk add --no-cache git curl build-base bash

# Install Air for live reloading
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

# Install Swag for generating Swagger documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Set the Go bin directory in PATH
ENV PATH=$PATH:/go/bin

# Create an app directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Expose port 8080
EXPOSE 8080

# Run Air
CMD ["air", "-c", ".air.toml"]
