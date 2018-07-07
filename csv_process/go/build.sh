#!/bin/bash
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.10 go build -v -ldflags "-s -w" -o app
