
#include <ndk.h>

static ngx_int_t ngx_http_l_set_host(ngx_http_request_t *r, ngx_str_t *val, ngx_http_variable_value_t *v);

static ndk_set_var_t ngx_http_l_set_host_filter = {
  NDK_SET_VAR_VALUE,
  (void *) ngx_http_l_set_host,
  1,
  NULL
};

static ngx_command_t ngx_http_l_commands[] = {
  {
    ngx_string ("set_host"),
    NGX_HTTP_LOC_CONF|NGX_CONF_TAKE1,
    ndk_set_var_value,
    0,
    0,
    &ngx_http_l_set_host_filter
  },
  ngx_null_command
};

static ngx_http_module_t ngx_http_l_module_ctx = {
  NULL,                         /* preconfiguration */
  NULL,                         /* postconfiguration */

  NULL,                         /* create main configuration */
  NULL,                         /* init main configuration */

  NULL,                         /* create server configuration */
  NULL,                         /* merge server configuration */

  NULL,   /* create location configuration */
  NULL     /*  merge location configuration */
};

ngx_module_t ngx_http_l_module = {
  NGX_MODULE_V1,
  &ngx_http_l_module_ctx,          /* module context */
  ngx_http_l_commands,             /* module directives */
  NGX_HTTP_MODULE,                 /* module type */
  NULL,                            /* init master */
  NULL,                            /* init module */
  NULL,                            /* init process */
  NULL,                            /* init thread */
  NULL,                            /* exit thread */
  NULL,                            /* exit process */
  NULL,                            /* exit master */
  NGX_MODULE_V1_PADDING
};

static ngx_int_t ngx_http_l_set_host(ngx_http_request_t *r, ngx_str_t *res, ngx_http_variable_value_t *v) {
  
  void *go_module = dlopen("/lab/ngx_http_l_module.a", RTLD_LAZY);
  if (!go_module) {
     //todo dlerror()
  }

  u_char* (*fun)(u_char *) = (u_char* (*)(u_char *)) dlsym(go_module, "LookupHost");
  
  ngx_str_t host = ngx_string(fun("test"));
  
  res->data = ngx_palloc(r->pool, host.len);
  if (res->data == NULL) {
    return NGX_ERROR;
  }

  ngx_memcpy(res->data, host.data, host.len);

  res->len = host.len;
  
  v->valid = 1;
  v->no_cacheable = 0;
  v->not_found = 0;
  
  return NGX_OK;
}


