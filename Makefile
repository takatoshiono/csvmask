VERSION=0.0.1
BUILD_LDFLAGS="-X main.version=$(VERSION)"

.PHONY: test
test:
	@go test -v -race ./...

.PHONY: build
build:
	@go build -ldflags=$(BUILD_LDFLAGS) ./...
