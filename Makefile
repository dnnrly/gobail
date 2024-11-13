GO111MODULE=on

CURL_BIN ?= curl
GO_BIN ?= go
GORELEASER_BIN ?= goreleaser

PUBLISH_PARAM?=
GO_MOD_PARAM?=-mod vendor
TMP_DIR?=./tmp

BASE_DIR=$(shell pwd)

NAME=gobail

export GO111MODULE=on
export GOPROXY=https://proxy.golang.org
export PATH := $(BASE_DIR)/bin:$(PATH)

help:   ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sort | sed -e 's/\\$$//' | sed -e 's/:.\+##/ --/'

.PHONY: clean
clean: ## remove build artifacts and release directories
	rm -f $(NAME)
	rm -rf dist/
	rm -rf cmd/$(NAME)/dist

.PHONY: deps
deps: build-deps test-deps ## ci target - install all runtime dependencies

.PHONY: test
test: ## run unit tests and format for human consumption
	$(GO_BIN) test -json ./... | tparse -all

.PHONY: acceptance-test
acceptance-test: ## run acceptance tests against the build gobailrm -rf ./test/tmp
	cd test && go build -cover -o testapp
	mkdir -p ./test/tmp/coverage
	cd test && GOCOVERDIR=tmp/coverage go test -timeout 20s -tags acceptance

.PHONY: ci-test
ci-test: ## ci target - run tests to generate coverage data
	$(GO_BIN) test -race -cover ./...

.PHONY: lint
lint: ## run linting
	golangci-lint run

.PHONY: coverage-report
coverage-report: ## collate the coverage data
	mkdir -p tmp/coverage
	go tool covdata textfmt -i=test/tmp/coverage,tmp/coverage -o ./coverage.txt
