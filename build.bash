#!/bin/bash
protoc -I./proto -I./deps --go_out=./go/proto --go-grpc_out=./go/proto --grpc-gateway_out=./go/proto --grpc-gateway_opt logtostderr=true --openapiv2_out ./swagger proto/*
