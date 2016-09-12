FROM debian:jessie

ARG GO_VERSION

RUN apt-get update
RUN apt-get install -y build-essential curl bash nano zlib1g-dev libpcre3-dev libssl-dev git socat
RUN curl -Os https://storage.googleapis.com/golang/go$GO_VERSION.linux-amd64.tar.gz && tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/lab

RUN mkdir /lab
WORKDIR /lab

ADD vendor.sh /lab
ARG NGINX_VERSION
ARG NDK_VERSION
RUN NGINX_VERSION=$NGINX_VERSION NDK_VERSION=$NDK_VERSION /lab/vendor.sh

COPY . /lab
RUN NGINX_VERSION=$NGINX_VERSION NDK_VERSION=$NDK_VERSION /lab/build.sh
