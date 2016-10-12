#!/usr/bin/env bash

WORKDIR=${1}

docker run --rm -i \
    --add-host "gocql-ec2.test:10.10.220.153" \
    --add-host "gocql-ec2.test:10.10.220.154" \
    -v "${WORKDIR}":/go -w /go \
    -e GOPATH=/go \
    golang:latest \
    go test --tags integration github.com/signal/gocql-ec2
