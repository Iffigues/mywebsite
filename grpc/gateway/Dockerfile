from golang:latest

RUN apt -y update
RUN apt -y upgrade
ADD ./gateway /gateway
workdir /gateway
RUN go mod init gateway
RUN go mod tidy
RUN go build
CMD ["./gateway"]
