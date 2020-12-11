FROM golang:1.15-alpine
WORKDIR /go/src/github.com/weblair/lair

RUN apk add git make docker-cli bash openssh-client

COPY ./ ./
RUN make clean
RUN make
RUN make install

CMD ["bash"]
