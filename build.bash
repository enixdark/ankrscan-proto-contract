#!/bin/bash
protoc -I./proto -I./deps --go_out=generated/go --go-grpc_out=generated/go --grpc-gateway_out=generated/go proto/*
