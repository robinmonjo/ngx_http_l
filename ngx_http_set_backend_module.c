
#include <ndk.h>

static ngx_int_t ngx_http_set_backend(ngx_http_request_t *r, ngx_str_t *val, ngx_http_variable_value_t *v);

static ndk_set_var_t ngx_http_set_backend_filter = {
  NDK_SET_VAR_VALUE,
  (void *) ngx_http_set_backend,
  1,
  NULL
};

static ngx_command_t ngx_http_set_backend_commands[] = {
  {
    ngx_string ("set_backend"),
    NGX_HTTP_LOC_CONF|NGX_CONF_TAKE1,
    ndk_set_var_value,
    0,
    0,
    &ngx_http_set_backend_filter
  },
  ngx_null_command
};

static ngx_http_module_t ngx_http_set_backend_module_ctx = { NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL };

ngx_module_t ngx_http_set_backend_module = {
  NGX_MODULE_V1,
  &ngx_http_set_backend_module_ctx,          /* module context */
  ngx_http_set_backend_commands,             /* module directives */
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

static ngx_int_t ngx_http_set_backend(ngx_http_request_t *r, ngx_str_t *res, ngx_http_variable_value_t *v) {
  
  void *go_module = dlopen("/lab/ngx_http_set_backend_module.a", RTLD_LAZY); //TODO: no hardcoded path
  if (!go_module) {
     //todo dlerror()
  }

  u_char* (*fun)(u_char *) = (u_char* (*)(u_char *)) dlsym(go_module, "LookupBackend");
  u_char* backend = fun(r->headers_in.host->value.data);
  
  ngx_str_t ngx_backend = { strlen(backend), backend };
  
  res->data = ngx_palloc(r->pool, ngx_backend.len);
  if (res->data == NULL) {
    return NGX_ERROR;
  }

  ngx_memcpy(res->data, ngx_backend.data, ngx_backend.len);

  res->len = ngx_backend.len;

  return NGX_OK;
}


