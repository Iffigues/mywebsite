FROM golang:latest

ENV PATH="$PATH:$(go env GOPATH)/bin"
RUN apt update -y
RUN apt upgrade -y
RUN apt install -y protobuf-compiler
RUN apt install -y golang-goprotobuf-dev
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
ADD ./chat.proto /proto/chat.proto
ADD ./chat.go /proto/chat.go
workdir /proto
RUN protoc --proto_path=. *.proto --go_out=. --go-grpc_out=.
