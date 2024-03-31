.PHONY: run build test clean

# Run the server
run:
	go run main.go

# Build the server
build:
	go build -o log-analyzer main.go

# Run tests
test:
	go test ./...

# Clean up build artifacts
clean:
	go clean
	rm -f log-analyzer
