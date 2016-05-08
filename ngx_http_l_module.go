// package name: ngx_http_l_module
package main

import "C"
import (
	"io/ioutil"
)

//export LookupHost
func LookupHost(r *C.char) *C.char {
	ioutil.WriteFile("/lab/build/nginx/logs/go.logs", []byte("salut ma gueule maudite"), 0644)
	return C.CString("ta racen la pute nerge")
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
