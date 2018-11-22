#/bin/bash

echo "Building..."


cd ../
export GOPATH=$(pwd)
cd ./src/picard/main
go build
# scp ...

