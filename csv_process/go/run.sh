#!/bin/bash
/usr/bin/time -v docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.10 ./app
