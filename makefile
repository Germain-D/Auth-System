.PHONY: install start build clean

# Default target
all: install start

# Install dependencies
install:
    @echo "Installing dependencies..."
    cd Auth-System-Front && npm install
    cd Auth-System-Back && go mod tidy

# Start development servers
start:
    @echo "Starting servers..."
    start cmd /c "cd Auth-System-Front && npm run dev"
    start cmd /c "cd Auth-System-Back && go run main.go"

# Build for production
build:
    @echo "Building for production..."
    cd Auth-System-Front && npm run build
    cd Auth-System-Back && go build -o dist/server.exe

# Clean build artifacts
clean:
    @echo "Cleaning..."
    cd Auth-System-Front && rm -rf .nuxt dist node_modules
    cd Auth-System-Back && rm -rf dist vendor