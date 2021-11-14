FROM golang:alpine

LABEL maintainer="Xuyuanteng <13883986114@163.com>"
WORKDIR $GOPATH/src/gin_docker
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
EXPOSE 9090
