FROM postgres:12
WORKDIR /root/go/src/github.com/weblair/lair

# Setup
RUN apt update
RUN apt install -y wget git make gcc
RUN mkdir -p /usr/local

# Setup Postgres
RUN echo "postgres" > /pgpw
RUN chown postgres:postgres /pgpw
RUN pg_createcluster 12 lair -- --auth=password --pwfile=/pgpw
RUN rm /pgpw

# Install Go
RUN wget https://golang.org/dl/go1.15.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
RUN rm go1.15.6.linux-amd64.tar.gz
ENV PATH $PATH:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin

# Install Lair
COPY ./ ./
RUN make clean
RUN make
RUN make install

CMD ["bash"]