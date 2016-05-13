#!/bin/bash

set -o nounset
set -o errexit
set -e

rm -rf build

#build go shared library
echo "building shared library"

CGO_CFLAGS="-I ./vendor/ngx_devel_kit-$NDK_VERSION/src" go build -o ngx_http_set_backend_module.a -buildmode=c-shared ngx_http_set_backend_module.go

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