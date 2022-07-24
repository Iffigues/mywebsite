FROM  golang:latest

ENV PATH="$PATH:$(go env GOPATH)/bin"
RUN apt -y update
RUN apt -y upgrade
RUN apt install -y fortune
RUN apt install -y fortunes-fr
RUN apt install -y fortunes-off
RUN apt-get install -y figlet
RUN apt-get install -y cowsay
RUN apt install -y toilet
RUN apt install -y rig
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

