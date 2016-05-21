#!/bin/bash

set -o nounset
set -o errexit

ROOT_DIR=`pwd`
NGX_VENDOR_DIR=$ROOT_DIR/src/ngx_http_set_backend/vendor
GO_VENDOR_DIR=$ROOT_DIR/src/backend/vendor


# downloading build dependencies into vendor
mkdir -p $NGX_VENDOR_DIR $GO_VENDOR_DIR
cd $NGX_VENDOR_DIR

# download ngx_devel_kit module
curl -s -L -O https://github.com/simpl/ngx_devel_kit/archive/v$NDK_VERSION.tar.gz
tar xzf  "v$NDK_VERSION.tar.gz"
rm -f "v$NDK_VERSION.tar.gz"

# download nginx
curl -s -L -O "http://nginx.org/download/nginx-$NGINX_VERSION.tar.gz"
tar xzf "nginx-$NGINX_VERSION.tar.gz"
rm -f "nginx-$NGINX_VERSION.tar.gz"

# download boltdb
cd $GO_VENDOR_DIR
git clone https://github.com/boltdb/bolt.git github.com/boltdb/bolt
git clone https://github.com/gorilla/mux.git github.com/gorilla/mux
git clone https://github.com/gorilla/context.git github.com/gorilla/context