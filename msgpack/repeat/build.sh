#!/usr/bin/env sh

rm pub sub
go build -o pub pub.go type.go
go build -o sub sub.go type.go