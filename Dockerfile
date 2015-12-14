FROM golang:1.4
MAINTAINER zaddone zaddone@qq.com
#ADD sources.list /etc/apt/sources.list 
RUN mkdir -p  "/data_sqlite" && chmod -R 777  "/data_sqlite"
VOLUME ["/data_sqlite"]
RUN go get github.com/beego/bee
RUN go get github.com/astaxie/beego
RUN go get github.com/garyburd/redigo/redis
RUN go get github.com/mattn/go-sqlite3
ADD . $GOPATH/src/github.com/zaddone/dimon
WORKDIR $GOPATH/src/github.com/zaddone/dimon
EXPOSE 80
CMD bee run 
