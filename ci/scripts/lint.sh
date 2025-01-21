#!/bin/bash -eux

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.0

pushd dp-topic-api
  make lint
popd
