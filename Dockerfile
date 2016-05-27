FROM alpine
RUN apk update && \
    apk add git && \
    apk add --no-cache -X http://dl-4.alpinelinux.org/alpine/edge/community go-tools && \
    mkdir -p /gopath/{bin,src/github.com/mwmahlberg/cayleybench,pkg}
ENV GOROOT=/usr/lib/go \
    GOPATH=/gopath \
    GOBIN=/gopath/bin \
    PATH=${PATH}:/goroot/bin:/gopath/bin
ADD . /gopath/src/github.com/mwmahlberg/cayleybench
WORKDIR /gopath/src/github.com/mwmahlberg/cayleybench
RUN go get "github.com/google/cayley" && \
    go get "github.com/pborman/uuid" && \
    go get "github.com/boltdb/bolt" && \
    go get "github.com/gogo/protobuf/proto" && \
    go get gopkg.in/mgo.v2 && \
    go get "github.com/lib/pq" && \
    go test -c && cp cayleybench.test /bin
ENTRYPOINT ["/bin/cayleybench.test", "-test.run=XXX" ,"-test.bench=.", "-test.benchmem","-sleep=5"]
CMD ["-test.cpu","1,2,4,6,8"]
