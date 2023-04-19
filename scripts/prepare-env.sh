#!/bin/bash
set -e

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.52.2
go install github.com/roblaszczak/go-cleanarch@latest