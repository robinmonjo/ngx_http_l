// package name: ngx_http_l_module
package main

// #include <nginx.h>
// #include <ngx_config.h>
// #include <ngx_core.h>
// #include <ngx_http.h>
import "C"
import (
	"io/ioutil"
)

//export HandleRequest
func HandleRequest(r *C.ngx_http_request_t) C.ngx_int_t {
	ioutil.WriteFile("/lab/build/nginx/logs/go.logs", []byte("salut ma gueule"), 0644)
	return C.NGX_DECLINED
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
