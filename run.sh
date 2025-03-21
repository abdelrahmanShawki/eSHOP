#!/bin/bash

SERVICE_NAME=$1
RELEASE_VERSION=$2

# Install necessary dependencies
sudo apt-get install -y protobuf-compiler golang-goprotobuf-dev

# Install Go plugins for Protocol Buffers and gRPC
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate Go code into the correct project structure
protoc --go_out=./${SERVICE_NAME}/internal/adapter/grpc/pb --go_opt=paths=source_relative \
  --go-grpc_out=./${SERVICE_NAME}/internal/adapter/grpc/pb --go-grpc_opt=paths=source_relative \
  ./proto/${SERVICE_NAME}.proto

# Navigate to the service directory and initialize Go module
cd ${SERVICE_NAME}
go mod init github.com/abdelrahmanShawki/eSHOP/${SERVICE_NAME} || true
go mod tidy

# Go back to the root directory
cd ../

# Configure Git
git config --global user.email "abdelrahman_shawki2017@outlook.com"
git config --global user.name "abdelrahmanShawki"

# Commit generated files and push the tag
git add . && git commit -am "proto update" || true
git tag -fa ${SERVICE_NAME}/${RELEASE_VERSION} -m "${SERVICE_NAME}/${RELEASE_VERSION}"
git push origin refs/tags/${SERVICE_NAME}/${RELEASE_VERSION}
