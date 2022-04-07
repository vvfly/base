#!/bin/sh

# cmd
set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=linux
go env -w GOOS=linux

# compile
go build -o main main.go

# reset
go env -w GOARCH=amd64
go env -w GOOS=windows
