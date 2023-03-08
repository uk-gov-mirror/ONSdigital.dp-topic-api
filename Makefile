BINPATH ?= build

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

BUILD_TIME=$(shell date +%s)
GIT_COMMIT=$(shell git rev-parse HEAD)
VERSION ?= $(shell git tag --points-at HEAD | grep ^v | head -n 1)

LDFLAGS = -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"

.PHONY: all
all: delimiter-AUDIT audit delimiter-LINTERS lint delimiter-UNIT-TESTS test delimiter-COMPONENT_TESTS test-component delimiter-FINISH ## Runs multiple targets, audit, lint, test and test-component

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

.PHONY: debug-run
debug-run:
	HUMAN_LOG=1 DEBUG=1 go run -tags 'debug' $(LDFLAGS) main.go

.PHONY: delimiter-%
delimiter-%:
	@echo '===================${GREEN} $* ${RESET}==================='

.PHONY: lint-local
lint-local:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	golangci-lint run ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -race -cover -tags 'production' ./...

.PHONY: test-component
test-component:
	go test -race -cover -coverprofile="coverage.txt" -coverpkg=github.com/ONSdigital/dp-topic-api/... -component

.PHONY: convey
convey:
	goconvey ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: help
help: ## Show help page for list of make targets
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
    }' $(MAKEFILE_LIST)
