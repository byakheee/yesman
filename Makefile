# メタ情報
NAME := yesman
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
	-X 'main.revision=$(REVISION)'

# 必要なツール類をセットアップする
## Setup
setup:
	go get github.com/Masterminds/glide
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help

# テストを実行する
## Run tests
test: deps
	go test $$(glide novendor)

# glideを使って依存パッケージをインストールする
## Install dependencies
deps: setup
	glide install

## Update dependencies
update: setup
	glide update

## Lint
lint: setup
	go vet $$(glide novendor)
	for pkg in $$(glide novendor -x); do \
		golint -set_exit_status $$pkg || exit $$?;\
	done

## Formmat source codes
fmt: setup
	goimports -w $$(glide nv -x)

## build binaries
build: cmd/yesman/main.go deps
	go build -ldflags "$(LDFLAGS)" -o bin/yesman $<

## build debug binaries
build-debug: cmd/yesman/main.go deps
	go build -tags=debug -ldflags "$(LDFLAGS)" -o bin/yesman $<

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: setup deps update test lint help