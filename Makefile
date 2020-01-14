GO111MODULE=on

CURL_BIN ?= curl
GO_BIN ?= go
GORELEASER_BIN ?= goreleaser

PUBLISH_PARAM?=
GO_MOD_PARAM?=-mod vendor
TMP_DIR?=./tmp

BASE_DIR=$(shell pwd)

NAME=paragraphical

export GO111MODULE=on
export GOPROXY=https://proxy.golang.org
export PATH := $(BASE_DIR)/bin:$(PATH)

install: deps

build:
	$(GO_BIN) build -v ./...

clean:
	rm -f $(NAME)
	rm -rf dist/
	rm -rf cmd/$(NAME)/dist

clean-deps:
	rm -rf ./bin
	rm -rf ./tmp
	rm -rf ./libexec
	rm -rf ./share

./bin/golangci-lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.17.1

./bin/goreleaser: ./bin ./tmp
	$(CURL_BIN) --fail -L -o ./tmp/goreleaser.tar.gz https://github.com/goreleaser/goreleaser/releases/download/v0.117.2/goreleaser_Linux_x86_64.tar.gz
	gunzip -f ./tmp/goreleaser.tar.gz
	tar -C ./bin -xvf ./tmp/goreleaser.tar

build-deps: ./bin/goreleaser

deps: build-deps

test:
	$(GO_BIN) test ./...

ci-test:
	$(GO_BIN) test -race -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	golangci-lint run

release: clean
	cd cmd/$(NAME) ; $(GORELEASER_BIN) $(PUBLISH_PARAM)

update:
	$(GO_BIN) get -u
	$(GO_BIN) mod tidy
	make test
	make install
	$(GO_BIN) mod tidy
