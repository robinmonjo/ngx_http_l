// package name: ngx_http_l_module
package main

import "C"

//export LookupHost
func LookupHost(r *C.char) *C.char {
	//TODO perform real lookup
	return C.CString("www.helloworld.nz")
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
