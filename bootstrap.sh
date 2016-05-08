#!/bin/bash

set -o nounset
set -o errexit

DIR=$(pwd)
BUILDDIR=$DIR/build
NGINX_DIR=nginx

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

# install nginx with the echo module and the ngx_devel_kit module
install_nginx () {
  if [ ! -d "vendor/nginx-$NGINX_VERSION" ]; then
    pushd vendor > /dev/null 2>&1
    # download echo module
    curl -s -L -O https://github.com/openresty/echo-nginx-module/archive/v$ECHO_VERSION.tar.gz
    tar xzf "v$ECHO_VERSION.tar.gz"
    rm -f "v$ECHO_VERSION.tar.gz"
    
    # download ngx_devel_kit module
    curl -s -L -O https://github.com/simpl/ngx_devel_kit/archive/v$NDK_VERSION.tar.gz
    tar xzf  "v$NDK_VERSION.tar.gz"
    rm -f "v$NDK_VERSION.tar.gz"
    
    ## download echo module (used for debugging)
    curl -s -L -O "http://nginx.org/download/nginx-$NGINX_VERSION.tar.gz"
    tar xzf "nginx-$NGINX_VERSION.tar.gz"
    rm -f "nginx-$NGINX_VERSION.tar.gz"
    
    pushd "nginx-$NGINX_VERSION" > /dev/null 2>&1
    ./configure                                             \
    --with-debug                                            \
    --prefix=$(pwd)/../../build/nginx                       \
    --conf-path=conf/nginx.conf                             \
    --error-log-path=logs/error.log                         \
    --http-log-path=logs/access.log                         \
    --add-module=../echo-nginx-module-$ECHO_VERSION         \
    --add-module=../ngx_devel_kit-$NDK_VERSION
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