GOOS=linux
GOARCH=amd64

.PHONY: build

GIT_COMMIT := $(shell git rev-list -1 HEAD)

build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-linkmode external -extldflags -static -X main.GitCommit=$(GIT_COMMIT)" -a main.go