#!/bin/bash
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.8.0/protoc-3.8.0-linux-x86_64.zip
unzip protoc-3.8.0-linux-x86_64.zip -d protoc3

rsync -ah ./protoc3/bin     /usr/local
rsync -ah ./protoc3/include /usr/local

chown     `whoami`     /usr/local/bin/protoc
chown -R  `whoami`     /usr/local/include/google

protoc --version
