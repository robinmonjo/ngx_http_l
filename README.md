# ngx_http_l

This is a work in progress. The goal is to develop a nginx plugin in Go (by making mandatory C code call Go code).

To be continued ...

Goals:

````
location / {
  set $gh_pages_host "";
  set $gh_pages_path "";

  access_by_lua_file /data/pages-lua/router.lua;

  proxy_set_header X-GitHub-Pages-Root $gh_pages_path;
  proxy_pass http://$gh_pages_host$request_uri;
}
````

resources: 
http://githubengineering.com/rearchitecting-github-pages/
http://www.nginxguts.com/2011/09/configuration-directives/#more-343

TODO:
- set a variable from a module with a directive
- load testing and dlopen performance documentation


### Hacking

`docker` must be installed and running

1. `make` - compile the module
2. `make test` - run integration tests

resources:
- http://blog.ralch.com/tutorial/golang-sharing-libraries/
- https://www.airpair.com/nginx/extending-nginx-tutorial

### Issues encountered

When calling a function from my shared library (written in Go), I sometime get locked forever on a futex during my request. To solve this I experimented **a lot**:
- tried to tweak my docker image in all possible ways
- used `ltrace` and `strace` to debug
- tried to use the dynamic module feature of nginx: https://www.nginx.com/blog/dynamic-modules-nginx-1-9-11/

I finally found [this issue](https://github.com/golang/go/issues/12873) on the Go github repository, that basically taught me that the Go runtime is loaded when the module is loaded by nginx and that Go library built with `buildmode=c-shared` should never get loaded before a `fork` (if the forked process intend to use the shared library). nginx workers are forked by the master process and they use the library. That was my problem. To solve it, and make the library calls work consistently I used `dlfcn` (`dlopen`, `dlsym`). This allow me to load dynamically the library in the workers (so after the `fork`). This probably has a performance impact, but I don't really care for now :)
