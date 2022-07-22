FROM golang:latest

RUN apt update -y
RUN apt upgrade -y
RUN apt install -y protobuf-compiler
RUN apt install -y golang-goprotobuf-dev

ADD ./chat.proto /proto/chat.proto
workdir /proto
RUN go mod init chat
