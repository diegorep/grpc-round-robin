#!/bin/bash
# Before using, you must install either the protoc binary (hard) or the protoc
# code generation tools in your language(s) of choice. Here I've left the
# examples for node, python and go

# Python installation:
# pip install grpcio
# pip install grpc-tools

# Node installation:
# npm install grpc-tools

# Python code-gen:
# Code-gen both the proto message and client/server stubs in python for all
# three services
python -m grpc_tools.protoc --proto_path=./proto/ --grpc_python_out=./python/services/ --python_out=./python/messages/ ones.proto
python -m grpc_tools.protoc --proto_path=./proto/ --grpc_python_out=./python/services/ --python_out=./python/messages/ twos.proto
#python -m grpc_tools.protoc --proto_path=./proto/ --grpc_python_out=./python/services/ --python_out=./python/messages/ threes.proto

# Node code-gen:
# Code-gen both the proto message and client/server stubs in NodeJS for all
# three services
grpc_tools_node_protoc --proto_path=./proto/ --grpc_out=./node/services/ --js_out=import_style=commonjs,binary:./node/messages/. --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` ones.proto
grpc_tools_node_protoc --proto_path=./proto/ --grpc_out=./node/services/ --js_out=import_style=commonjs,binary:./node/messages/. --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` twos.proto
#grpc_tools_node_protoc --proto_path=./proto/ --grpc_out=./node/services/ --js_out=import_style=commonjs,binary:./node/messages/. --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` threes.proto
