#!/bin/sh

protoc -I . user.proto --go_out=plugins=grpc:../src