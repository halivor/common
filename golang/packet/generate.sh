#!/bin/bash

protoc -I. -I/home/dev/gopath/src/github.com/halivor/common/golang/packet/ \
    --go_out=plugins=grpc:. ./packet.proto
mv github.com/halivor/common/golang/packet/packet.pb.go .
rm -rf ./github.com/
