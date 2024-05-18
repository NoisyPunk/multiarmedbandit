BIN := "./bin"
DOCKER_IMG="rotator:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build-rotator:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd

run-rotator: build-rotator
	$(BIN)/rotator -config ./configs/rotator_config.yaml

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

generate:
	go generate ./...

test:
	go test -race ./internal/...

integration-tests:
	docker-compose up -d --force-recreate && \
	docker-compose -f docker-compose.integration_tests.yml up --force-recreate && \
	docker-compose -f docker-compose.integration_tests.yml down && \
	docker-compose down

install-lint-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: install-lint-deps
	golangci-lint run ./...

up:
	docker-compose up --force-recreate

down:
	docker-compose down

.PHONY: build-rotator run-rotator build-img run-img version test lint generate
