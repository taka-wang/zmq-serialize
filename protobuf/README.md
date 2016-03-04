# Protocol buffer

## Setup

#### @node.js package

`npm install protocol-buffers`

#### Golang

```bash
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go

# generate go code from proto v3
protoc --plugin=protoc-gen-go  --go_out=. type.proto
```

#### C

```bash
# generate c code from proto v2
protoc-c --c_out=. type2.proto
```

#### Warning
***protoc-c will rename keyword "register" to "register_"***
