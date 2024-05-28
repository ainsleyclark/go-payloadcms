setup: # Setup dependencies
	go mod tidy
	go generate ./...
	cd dev && npm install
.PHONY: setup

lint: # Run linter and fix all issues
	golangci-lint run --fix ./...
.PHONY: lint

format: # Format all code
	go fmt ./...
.PHONY: format

excluded := grep -v /gen/ | grep -v /fakes/ | grep -v /tests/ | grep -v /dev/

test: # Run all tests with race detection and coverage
	go test -race -coverprofile=coverage.out -covermode=atomic $$(go list ./... | $(excluded)) ./...
.PHONY: test

generate: # Runs go generate
	go generate ./...
.PHONY: generate

cover: # Run all the tests and opens the coverage report
	go tool cover -html=coverage.out
.PHONY: cover

all: # Make format, lint and test
	$(MAKE) format
	$(MAKE) lint
	$(MAKE) test
.PHONY: all

todo: # Show to-do items per file
	$(Q) grep \
		--exclude=Makefile.util \
		--exclude-dir=vendor \
		--exclude-dir=.vercel \
		--exclude-dir=.gen \
		--exclude-dir=.idea \
		--exclude-dir=public \
		--exclude-dir=node_modules \
		--exclude-dir=archetypes \
		--exclude-dir=.git \
		--text \
		--color \
		-nRo \
		-E '\S*[^\.]TODO.*' \
		.
.PHONY: todo

lines: # Show line count of Go code
	find . -name '*.go' | xargs wc -l
.PHONY: lines

help: # Display this help
	$(Q) awk 'BEGIN {FS = ":.*#"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?#/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help
