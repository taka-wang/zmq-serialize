#!/usr/bin/env sh

go build -o pub pub.go type.go
go build -o sub sub.go type.go