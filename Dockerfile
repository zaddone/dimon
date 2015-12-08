FROM golang:1.4
MAINTAINER zaddone zaddone@qq.com
RUN go get github.com/beego/bee
ADD . $GOPATH/src/github.com/zaddone/dimon
WORKDIR $GOPATH/src/github.com/zaddone/dimon
EXPOSE 8080
CMD bee run
