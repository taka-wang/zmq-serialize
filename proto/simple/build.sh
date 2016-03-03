#!/usr/bin/env sh

rm pubc subc
gcc type2.pb-c.h type2.pb-c.c pub.c -lzmq -lczmq -lprotobuf-c  -o pubc
gcc type2.pb-c.h type2.pb-c.c sub.c -lzmq -lczmq -lprotobuf-c  -o subc