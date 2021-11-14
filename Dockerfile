FROM golang:alpine

LABEL maintainer="Xuyuanteng <13883986114@163.com>"
WORKDIR $GOPATH/src/gin_docker
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
RUN go build -o gin_docker .
EXPOSE 8080
ENTRYPOINT  ["./gin_docker"]
