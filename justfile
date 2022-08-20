#!/usr/bin/env just --justfile

run file="README.md":
 go run . {{file}}

update:
  go get -u
  go mod tidy -v