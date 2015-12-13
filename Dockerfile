FROM ubuntu:14.04
MAINTAINER zaddone zaddone@qq.com
ADD sources.list /etc/apt/sources.list 
RUN apt-get update  
RUN apt-get -y install golang git gcc && apt-get clean
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
RUN go get github.com/beego/bee
RUN go get github.com/astaxie/beego
RUN go get github.com/nfnt/resize
#RUN go get github.com/zaddone/dimon
ADD . $GOPATH/src/github.com/zaddone/dimon
WORKDIR $GOPATH/src/github.com/zaddone/dimon
EXPOSE 8080
CMD bee run 
