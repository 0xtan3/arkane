# Run the application
run:
	@go run ./cmd/main.go

# Build the application
build:
	@go build -o arkane ./cmd/main.go

# Format the code
fmt:
	@go fmt ./...

# Tidy up the dependencies
tidy:
	@go mod tidy

# Clean the build
clean:
	@rm -f arkane

