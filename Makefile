include .env
export

openapi_generate:
	@./scripts/openapi-http.sh bot internal/bot/ports ports

fmt:
	goimports -l -w internal/

lint:
	@go-cleanarch
	@./scripts/lint.sh common
	@./scripts/lint.sh bot

prepare_env:
	@./scripts/prepare-env.sh

clean:
	rm -rf bin/*

test:
	@./scripts/test.sh bot .test.env

build:
	CGO_ENABLED=0 go build -o bin/bot-runner cmd/bot/main.go

build-dev: clean fmt lint test
	CGO_ENABLED=1 go build -race -o bin/bot-runner cmd/bot/main.go

c4:
	cd tools/c4 && go mod tidy && sh generate.sh

.NOTPARALLEL:
.PHONY: openapi_generate fmt lint prepare_env clean test build build-dev c4

