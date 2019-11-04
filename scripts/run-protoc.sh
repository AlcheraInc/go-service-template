#!/bin/bash

workspace=$(pwd)
proto_compiler_path=${1:-"$(which protoc)"}
grpc_go_plugin_path=${2:-"${workspace}/bin/protoc-gen-go"}

mkdir -p "${workspace}/src/echo_v1"

${proto_compiler_path} service.proto \
    --go_out="${workspace}/src/echo_v1"  \
    --plugin="protoc-gen-go=${grpc_go_plugin_path}"

${proto_compiler_path} service.proto \
    --go_out="plugins=grpc:${workspace}/src/echo_v1" \
    --plugin="protoc-gen-go=${grpc_go_plugin_path}"
