FROM golang:1.8

MAINTAINER 1ambda

COPY . /go/src/github.com/a-trium/gipeline/server-gateway
WORKDIR /go/src/github.com/a-trium/gipeline/server-gateway
RUN make install build

EXPOSE 10001

CMD ["make", "run"]







