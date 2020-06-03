FROM golang:1.12-alpine
WORKDIR /root

RUN apk add git make docker-cli bash openssh-client

RUN mkdir -p ./temp/lair/
COPY ./ ./temp/lair/
RUN cd ./temp/lair/ && make clean
RUN cd ./temp/lair/ && make
RUN cd ./temp/lair/ && make install
RUN rm -rf ./temp/

CMD ["bash"]
