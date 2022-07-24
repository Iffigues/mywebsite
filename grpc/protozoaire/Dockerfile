FROM golang:latest

ENV PATH="$PATH:$(go env GOPATH)/bin"
RUN apt update -y
RUN apt upgrade -y
RUN apt install -y protobuf-compiler
RUN apt install -y golang-goprotobuf-dev
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
ADD ./chat.proto /proto/chat.proto
ADD ./chat /proto/chat
workdir /proto/chat
RUN go mod init chat
RUN go mod tidy
RUN go get
RUN go build
WORKDIR /proto
RUN protoc --proto_path=. *.proto --go_out=. --go-grpc_out=.
