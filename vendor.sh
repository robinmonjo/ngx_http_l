#!/bin/bash

set -o nounset
set -o errexit

VENDOR_DIR=./src/set_backend/vendor

# downloading build dependencies into vendor
mkdir -p $VENDOR_DIR
cd $VENDOR_DIR

# download ngx_devel_kit module
curl -s -L -O https://github.com/simpl/ngx_devel_kit/archive/v$NDK_VERSION.tar.gz
tar xzf  "v$NDK_VERSION.tar.gz"
rm -f "v$NDK_VERSION.tar.gz"

# download nginx
curl -s -L -O "http://nginx.org/download/nginx-$NGINX_VERSION.tar.gz"
tar xzf "nginx-$NGINX_VERSION.tar.gz"
rm -f "nginx-$NGINX_VERSION.tar.gz"