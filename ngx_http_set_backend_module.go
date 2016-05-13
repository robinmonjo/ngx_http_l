// package name: ngx_http_set_backend_module
package main

import (
	"C"
)

//export LookupBackend
func LookupBackend(_host *C.char) *C.char {
	host := C.GoString(_host)

	//TODO perform lookup
	host = "www.google.com"

	return C.CString(host)
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
