#!/usr/bin/env sh

rm pubc subc pub sub

# build c
gcc type2.pb-c.h type2.pb-c.c pub.c -lzmq -lczmq -lprotobuf-c  -o pubc
gcc type2.pb-c.h type2.pb-c.c sub.c -lzmq -lczmq -lprotobuf-c  -o subc

# build golang
go build -o pub pub.go type.pb.go
go build -o sub sub.go type.pb.go