#!/usr/bin/env just --justfile

run:
 go run .

update:
  go get -u
  go mod tidy -v