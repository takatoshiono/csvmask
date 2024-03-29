VERSION := 0.1.2
BUILD_LDFLAGS := "-X main.version=$(VERSION)"

.PHONY: test
test:
	@go test -v -race ./...

.PHONY: coverage
coverage:
	@go test -v -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: build
build:
	@go build -ldflags=$(BUILD_LDFLAGS) ./...

.PHONY: install
install:
	@go install -ldflags=$(BUILD_LDFLAGS) ./...

.PHONY: devel-deps
devel-deps:
	@go get \
		github.com/Songmu/goxz/cmd/goxz \
		github.com/tcnksm/ghr

.PHONY: crossbuild
crossbuild:
	@goxz -build-ldflags=$(BUILD_LDFLAGS) \
		-pv $(VERSION) -os=darwin,linux,windows -arch=386,amd64 \
		-d ./pkg/v$(VERSION)

.PHONY: upload
upload:
	@ghr -u takatoshiono v$(VERSION) ./pkg/v$(VERSION)
