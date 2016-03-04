#!/usr/bin/env sh

rm pubc subc pub sub

# build golang
go build -o pub pub.go type.go
go build -o sub sub.go type.go