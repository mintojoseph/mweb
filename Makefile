GOOS=linux
GOARCH=amd64

.PHONY: build

GIT_COMMIT := $(shell git rev-list -1 HEAD)

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.GitCommit=$(GIT_COMMIT)" .