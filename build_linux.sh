#!/bin/bash

GOOS=linux GOARCH=amd64 /usr/local/go/bin/go build -ldflags '-w -s' -o runner
