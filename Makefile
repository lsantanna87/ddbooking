THIS_FILE := $(lastword $(MAKEFILE_LIST))
SERVICENAME=ddbooking
SHELL := /bin/bash

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: ci
ci: ## Simulates CI.
	@$(MAKE) -f $(THIS_FILE) lint
	@$(MAKE) -f $(THIS_FILE) test-with-coverage
	@$(MAKE) -f $(THIS_FILE) build
	@$(MAKE) -f $(THIS_FILE) cleanup

.PHONY: build
build: ## Build ddbooking Binary
	go build

.PHONY: cleanup
cleanup: ## Deletes Temp Files.
	rm -f ./coverage.txt
	rm -f ./coverage.html

.PHONY: lint
lint: ## Runs linter
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run --skip-files='(test)' --timeout 2m0s

test: ## Runs tests
	echo 'mode: atomic' > ./coverage.txt
	SERVICENAME=$(SERVICENAME) CONFIGDIR=$(CONFIGDIR) go test -covermode=atomic -coverprofile=./coverage.txt -v -p 1 -race -timeout=30s ./...
	./run-static-analysis.sh

test-with-coverage: ## Runs tests with coverage
	@$(MAKE) -f $(THIS_FILE) test
	go tool cover -html=./coverage.txt -o ./coverage.html
	go tool cover -func=./coverage.txt

.DEFAULT: help
