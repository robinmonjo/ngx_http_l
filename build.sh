#!/bin/bash

set -o nounset
set -o errexit
set -e

# build go shared library
ROOT_DIR=`pwd`
BUILD_DIR=$ROOT_DIR/build
mkdir -p $BUILD_DIR
cd src/ngx_http_set_backend

echo "building shared library"
CGO_CFLAGS="-I ./vendor/ngx_devel_kit-$NDK_VERSION/src" go build -o $BUILD_DIR/ngx_http_set_backend_module.a -buildmode=c-shared ngx_http_set_backend_module.go

# build nginx
cd vendor/nginx-$NGINX_VERSION
CFLAGS="-g -O0" ./configure           											\
    --with-debug                      											\
    --prefix=$BUILD_DIR/nginx 											        \
    --conf-path=conf/nginx.conf       											\
    --error-log-path=logs/error.log   											\
    --http-log-path=logs/access.log   											\
    --add-module=../ngx_devel_kit-$NDK_VERSION							\
    --add-module=../../
make
make install

ln -sf $ROOT_DIR/nginx.conf $BUILD_DIR/nginx/conf/nginx.conf

# build backends_store
echo "building backends_store"
cd $ROOT_DIR/src/backends_store
go build -o $BUILD_DIR/backends_store