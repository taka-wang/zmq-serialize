# Protocol buffer

## Setup

#### @node.js package

`npm install protocol-buffers`

#### @golang

```
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go

./protoc --plugin=protoc-gen-go  --go_out=. type.proto
```

#### @protoc-c
```
protoc-c --c_out=. type2.proto # code gen
```
