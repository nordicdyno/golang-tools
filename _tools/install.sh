#!/usr/bin/env bash
set -e
cd _tools
set -x
go run ./ls-imports/main.go -u -f tools.go | xargs -tI % go install -v %
