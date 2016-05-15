
default:	build

clean:
	rm -rf Makefile objs

build:
	$(MAKE) -f objs/Makefile

install:
	$(MAKE) -f objs/Makefile install

modules:
	$(MAKE) -f objs/Makefile modules

upgrade:
	/Users/robinmonjo/code/infra/ngx_http_l/build/nginx/sbin/nginx -t

	kill -USR2 `cat /Users/robinmonjo/code/infra/ngx_http_l/build/nginx/logs/nginx.pid`
	sleep 1
	test -f /Users/robinmonjo/code/infra/ngx_http_l/build/nginx/logs/nginx.pid.oldbin

	kill -QUIT `cat /Users/robinmonjo/code/infra/ngx_http_l/build/nginx/logs/nginx.pid.oldbin`
