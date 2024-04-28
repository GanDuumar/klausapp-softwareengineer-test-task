#/bin/sh

# This is for compiling grpc and protopuf files
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\ratings.proto
