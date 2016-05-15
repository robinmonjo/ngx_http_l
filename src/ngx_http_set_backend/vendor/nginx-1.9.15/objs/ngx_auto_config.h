#define NGX_CONFIGURE " --with-debug --prefix=/Users/robinmonjo/code/infra/ngx_http_l/build/nginx --conf-path=conf/nginx.conf --error-log-path=logs/error.log --http-log-path=logs/access.log --add-module=../ngx_devel_kit-0.3.0rc1 --add-module=../../"

#ifndef NGX_DEBUG
#define NGX_DEBUG  1
#endif


#ifndef NGX_HAVE_GCC_ATOMIC
#define NGX_HAVE_GCC_ATOMIC  1
#endif


#ifndef NGX_HAVE_C99_VARIADIC_MACROS
#define NGX_HAVE_C99_VARIADIC_MACROS  1
#endif


#ifndef NGX_HAVE_GCC_VARIADIC_MACROS
#define NGX_HAVE_GCC_VARIADIC_MACROS  1
#endif


#ifndef NGX_HAVE_GCC_BSWAP64
#define NGX_HAVE_GCC_BSWAP64  1
#endif


#ifndef NGX_HAVE_KQUEUE
#define NGX_HAVE_KQUEUE  1
#endif


#ifndef NGX_HAVE_CLEAR_EVENT
#define NGX_HAVE_CLEAR_EVENT  1
#endif


#ifndef NGX_HAVE_TIMER_EVENT
#define NGX_HAVE_TIMER_EVENT  1
#endif


#ifndef NGX_HAVE_SENDFILE
#define NGX_HAVE_SENDFILE  1
#endif


#ifndef NGX_DARWIN_ATOMIC
#define NGX_DARWIN_ATOMIC  1
#endif


#ifndef NGX_HAVE_NONALIGNED
#define NGX_HAVE_NONALIGNED  1
#endif


#ifndef NGX_CPU_CACHE_LINE
#define NGX_CPU_CACHE_LINE  64
#endif


#define NGX_KQUEUE_UDATA_T  (void *)


#ifndef NGX_HAVE_F_NOCACHE
#define NGX_HAVE_F_NOCACHE  1
#endif


#ifndef NGX_HAVE_STATFS
#define NGX_HAVE_STATFS  1
#endif


#ifndef NGX_HAVE_STATVFS
#define NGX_HAVE_STATVFS  1
#endif


#ifndef NGX_HAVE_DLOPEN
#define NGX_HAVE_DLOPEN  1
#endif


#ifndef NGX_HAVE_SCHED_YIELD
#define NGX_HAVE_SCHED_YIELD  1
#endif


#ifndef NGX_HAVE_REUSEPORT
#define NGX_HAVE_REUSEPORT  1
#endif


#ifndef NGX_HAVE_IP_RECVDSTADDR
#define NGX_HAVE_IP_RECVDSTADDR  1
#endif


#ifndef NGX_HAVE_IP_PKTINFO
#define NGX_HAVE_IP_PKTINFO  1
#endif


#ifndef NGX_HAVE_TCP_FASTOPEN
#define NGX_HAVE_TCP_FASTOPEN  1
#endif


#ifndef NGX_HAVE_UNIX_DOMAIN
#define NGX_HAVE_UNIX_DOMAIN  1
#endif


#ifndef NGX_PTR_SIZE
#define NGX_PTR_SIZE  8
#endif


#ifndef NGX_SIG_ATOMIC_T_SIZE
#define NGX_SIG_ATOMIC_T_SIZE  4
#endif


#ifndef NGX_HAVE_LITTLE_ENDIAN
#define NGX_HAVE_LITTLE_ENDIAN  1
#endif


#ifndef NGX_MAX_SIZE_T_VALUE
#define NGX_MAX_SIZE_T_VALUE  9223372036854775807LL
#endif


#ifndef NGX_SIZE_T_LEN
#define NGX_SIZE_T_LEN  (sizeof("-9223372036854775808") - 1)
#endif


#ifndef NGX_MAX_OFF_T_VALUE
#define NGX_MAX_OFF_T_VALUE  9223372036854775807LL
#endif


#ifndef NGX_OFF_T_LEN
#define NGX_OFF_T_LEN  (sizeof("-9223372036854775808") - 1)
#endif


#ifndef NGX_TIME_T_SIZE
#define NGX_TIME_T_SIZE  8
#endif


#ifndef NGX_TIME_T_LEN
#define NGX_TIME_T_LEN  (sizeof("-9223372036854775808") - 1)
#endif


#ifndef NGX_MAX_TIME_T_VALUE
#define NGX_MAX_TIME_T_VALUE  9223372036854775807LL
#endif


#ifndef NGX_HAVE_PREAD
#define NGX_HAVE_PREAD  1
#endif


#ifndef NGX_HAVE_PWRITE
#define NGX_HAVE_PWRITE  1
#endif


#ifndef NGX_SYS_NERR
#define NGX_SYS_NERR  107
#endif


#ifndef NGX_HAVE_LOCALTIME_R
#define NGX_HAVE_LOCALTIME_R  1
#endif


#ifndef NGX_HAVE_POSIX_MEMALIGN
#define NGX_HAVE_POSIX_MEMALIGN  1
#endif


#ifndef NGX_HAVE_MAP_ANON
#define NGX_HAVE_MAP_ANON  1
#endif


#ifndef NGX_HAVE_SYSVSHM
#define NGX_HAVE_SYSVSHM  1
#endif


#ifndef NGX_HAVE_MSGHDR_MSG_CONTROL
#define NGX_HAVE_MSGHDR_MSG_CONTROL  1
#endif


#ifndef NGX_HAVE_FIONBIO
#define NGX_HAVE_FIONBIO  1
#endif


#ifndef NGX_HAVE_GMTOFF
#define NGX_HAVE_GMTOFF  1
#endif


#ifndef NGX_HAVE_D_NAMLEN
#define NGX_HAVE_D_NAMLEN  1
#endif


#ifndef NGX_HAVE_D_TYPE
#define NGX_HAVE_D_TYPE  1
#endif


#ifndef NGX_HAVE_SC_NPROCESSORS_ONLN
#define NGX_HAVE_SC_NPROCESSORS_ONLN  1
#endif


#ifndef NGX_HAVE_OPENAT
#define NGX_HAVE_OPENAT  1
#endif


#ifndef NGX_HAVE_GETADDRINFO
#define NGX_HAVE_GETADDRINFO  1
#endif


#ifndef NGX_HTTP_CACHE
#define NGX_HTTP_CACHE  1
#endif


#ifndef NGX_HTTP_GZIP
#define NGX_HTTP_GZIP  1
#endif


#ifndef NGX_HTTP_SSI
#define NGX_HTTP_SSI  1
#endif


#ifndef NGX_CRYPT
#define NGX_CRYPT  1
#endif


#ifndef NGX_HTTP_X_FORWARDED_FOR
#define NGX_HTTP_X_FORWARDED_FOR  1
#endif


#ifndef NGX_HTTP_X_FORWARDED_FOR
#define NGX_HTTP_X_FORWARDED_FOR  1
#endif


#ifndef NGX_HTTP_UPSTREAM_ZONE
#define NGX_HTTP_UPSTREAM_ZONE  1
#endif


#ifndef NDK
#define NDK  1
#endif


#ifndef NDK_SET_VAR
#define NDK_SET_VAR  1
#endif


#ifndef NGX_PCRE
#define NGX_PCRE  1
#endif


#ifndef NGX_HAVE_PCRE_JIT
#define NGX_HAVE_PCRE_JIT  1
#endif


#ifndef NGX_ZLIB
#define NGX_ZLIB  1
#endif


#ifndef NGX_PREFIX
#define NGX_PREFIX  "/Users/robinmonjo/code/infra/ngx_http_l/build/nginx/"
#endif


#ifndef NGX_CONF_PREFIX
#define NGX_CONF_PREFIX  "conf/"
#endif


#ifndef NGX_SBIN_PATH
#define NGX_SBIN_PATH  "sbin/nginx"
#endif


#ifndef NGX_CONF_PATH
#define NGX_CONF_PATH  "conf/nginx.conf"
#endif


#ifndef NGX_PID_PATH
#define NGX_PID_PATH  "logs/nginx.pid"
#endif


#ifndef NGX_LOCK_PATH
#define NGX_LOCK_PATH  "logs/nginx.lock"
#endif


#ifndef NGX_ERROR_LOG_PATH
#define NGX_ERROR_LOG_PATH  "logs/error.log"
#endif


#ifndef NGX_HTTP_LOG_PATH
#define NGX_HTTP_LOG_PATH  "logs/access.log"
#endif


#ifndef NGX_HTTP_CLIENT_TEMP_PATH
#define NGX_HTTP_CLIENT_TEMP_PATH  "client_body_temp"
#endif


#ifndef NGX_HTTP_PROXY_TEMP_PATH
#define NGX_HTTP_PROXY_TEMP_PATH  "proxy_temp"
#endif


#ifndef NGX_HTTP_FASTCGI_TEMP_PATH
#define NGX_HTTP_FASTCGI_TEMP_PATH  "fastcgi_temp"
#endif


#ifndef NGX_HTTP_UWSGI_TEMP_PATH
#define NGX_HTTP_UWSGI_TEMP_PATH  "uwsgi_temp"
#endif


#ifndef NGX_HTTP_SCGI_TEMP_PATH
#define NGX_HTTP_SCGI_TEMP_PATH  "scgi_temp"
#endif


#ifndef NGX_SUPPRESS_WARN
#define NGX_SUPPRESS_WARN  1
#endif


#ifndef NGX_SMP
#define NGX_SMP  1
#endif


#ifndef NGX_USER
#define NGX_USER  "nobody"
#endif


#ifndef NGX_GROUP
#define NGX_GROUP  "nobody"
#endif

