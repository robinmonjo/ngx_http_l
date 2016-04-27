#!/bin/bash

set -o nounset
set -o errexit
set -e

rm -rf build

#build go shared library
echo "building shared library"

rm -f ngx_http_l_module.a ngx_http_l_module.h

CGO_CFLAGS="-I /lab/vendor/nginx-$NGINX_VERSION/src/core \
	-I /lab/vendor/nginx-$NGINX_VERSION/src/event \
	-I /lab/vendor/nginx-$NGINX_VERSION/src/event/modules \
	-I /lab/vendor/nginx-$NGINX_VERSION/src/os/unix \
	-I /lab/vendor/nginx-$NGINX_VERSION/usr/local/include \
	-I /lab/vendor/nginx-$NGINX_VERSION/objs \
	-I /lab/vendor/nginx-$NGINX_VERSION/src/http \
	-I /lab/vendor/nginx-$NGINX_VERSION/src/http/modules \
	-I /lab/vendor/nginx-$NGINX_VERSION/src/mail \
	-I /lab/vendor/nginx-$NGINX_VERSION/src/stream" go build -o ngx_http_l_module.a -buildmode=c-shared ngx_http_l_module.go

#compile nginx
pushd vendor > /dev/null 2>&1
pushd nginx-$NGINX_VERSION > /dev/null 2>&1
CFLAGS="-g -O0" ./configure           \
    --with-debug                      \
    --prefix=$(pwd)/../../build/nginx \
    --conf-path=conf/nginx.conf       \
    --error-log-path=logs/error.log   \
    --http-log-path=logs/access.log   \
    --add-module=../../ \
		--with-http_ssl_module 
make
make install
popd > /dev/null 2>&1
popd > /dev/null 2>&1

ln -sf $(pwd)/nginx.conf $(pwd)/build/nginx/conf/nginx.conf
chown -R nobody $(pwd)/build/nginx/logs