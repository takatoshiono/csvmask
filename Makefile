VERSION := 0.0.1
BUILD_LDFLAGS := "-X main.version=$(VERSION)"

.PHONY: test
test:
	@go test -v -race ./...

.PHONY: build
build:
	@go build -ldflags=$(BUILD_LDFLAGS) ./...

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
	@ghr v$(VERSION) ./pkg/v$(VERSION)
