BINPATH ?= build

BUILD_TIME=$(shell date +%s)
GIT_COMMIT=$(shell git rev-parse HEAD)
VERSION ?= $(shell git tag --points-at HEAD | grep ^v | head -n 1)

LDFLAGS = -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"

.PHONY: all
all: audit test build

.PHONY: audit
audit:
	go list -m all | nancy sleuth
	
.PHONY: build
build:
	go build -tags 'production' $(LDFLAGS) -o $(BINPATH)/dp-topic-api

.PHONY: debug
debug:
	go build -tags 'debug' $(LDFLAGS) -o $(BINPATH)/dp-topic-api
	HUMAN_LOG=1 DEBUG=1 $(BINPATH)/dp-topic-api

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: test-component
test-component:
	go test -race -cover -coverprofile="coverage.txt" -coverpkg=github.com/ONSdigital/dp-topic-api/... -component

.PHONY: convey
convey:
	goconvey ./...

