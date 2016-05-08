#!/bin/bash

set -o nounset
set -o errexit
set -e

rm -rf build

#build go shared library
echo "building shared library"

rm -f ngx_http_l_module.a ngx_http_l_module.h

CGO_CFLAGS="-I ./vendor/nginx-$NGINX_VERSION/src/core \
	-I ./vendor/nginx-$NGINX_VERSION/src/event \
	-I ./vendor/nginx-$NGINX_VERSION/src/event/modules \
	-I ./vendor/nginx-$NGINX_VERSION/src/os/unix \
	-I ./vendor/nginx-$NGINX_VERSION/usr/local/include \
	-I ./vendor/nginx-$NGINX_VERSION/objs \
	-I ./vendor/nginx-$NGINX_VERSION/src/http \
	-I ./vendor/nginx-$NGINX_VERSION/src/http/modules \
	-I ./vendor/nginx-$NGINX_VERSION/src/mail \
	-I ./vendor/nginx-$NGINX_VERSION/src/stream" go build -o ngx_http_l_module.a -buildmode=c-shared ngx_http_l_module.go

#compile nginx with the echo module
pushd vendor > /dev/null 2>&1
pushd nginx-$NGINX_VERSION > /dev/null 2>&1
CFLAGS="-g -O0" ./configure           											\
    --with-debug                      											\
    --prefix=$(pwd)/../../build/nginx 											\
    --conf-path=conf/nginx.conf       											\
    --error-log-path=logs/error.log   											\
    --http-log-path=logs/access.log   											\
		--add-module=../echo-nginx-module-$ECHO_VERSION  				\
    --add-module=../ngx_devel_kit-$NDK_VERSION							\
    --add-module=../../
make
make install
popd > /dev/null 2>&1
popd > /dev/null 2>&1

ln -sf $(pwd)/nginx.conf $(pwd)/build/nginx/conf/nginx.conf
chown -R nobody $(pwd)/build/nginx/logs