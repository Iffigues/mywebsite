from golang:latest

RUN apt -y update
RUN apt -u upgrade
ADD ./goclient /goclient
RUN mkdir /build
workdir /goclient
RUN go mod init goclient
RUN go mod tidy
RUN go build
CMD ["./goclient"]
