#!/bin/bash

set -o nounset
set -o errexit

DIR=$(pwd)
BUILDDIR=$DIR/build
NGINX_DIR=nginx
NGINX_VERSION=$NGINX_VERSION
ECHO_MODULE_VERSION=0.59rc1

clean () {
  rm -rf build vendor
}

setup_local_directories () {
  if [ ! -d $BUILDDIR ]; then
    mkdir $BUILDDIR > /dev/null 2>&1
    mkdir $BUILDDIR/$NGINX_DIR > /dev/null 2>&1
  fi

  if [ ! -d "vendor" ]; then
    mkdir vendor > /dev/null 2>&1
  fi
}

install_nginx () {
  if [ ! -d "vendor/nginx-$NGINX_VERSION" ]; then
    pushd vendor > /dev/null 2>&1
    curl -s -L -O https://github.com/openresty/echo-nginx-module/archive/v$ECHO_MODULE_VERSION.tar.gz
    tar xzf "v$ECHO_MODULE_VERSION.tar.gz"
    mv echo-nginx-module-$ECHO_MODULE_VERSION echo-nginx-module
    curl -s -L -O "http://nginx.org/download/nginx-$NGINX_VERSION.tar.gz"
    tar xzf "nginx-$NGINX_VERSION.tar.gz"
    pushd "nginx-$NGINX_VERSION" > /dev/null 2>&1
    ./configure                       \
    --with-debug                      \
    --prefix=$(pwd)/../../build/nginx \
    --conf-path=conf/nginx.conf       \
    --error-log-path=logs/error.log   \
    --http-log-path=logs/access.log   \
    --add-module=../echo-nginx-module
    make
    make install
    popd > /dev/null 2>&1
    popd > /dev/null 2>&1
  else
    printf "NGINX already installed\n"
  fi
}

if [[ "$#" -eq 1 ]]; then
  if [[ "$1" == "clean" ]]; then
    clean
  else
    echo "clean is the only option"
  fi
else
  setup_local_directories
  install_nginx
fi