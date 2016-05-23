// package name: ngx_http_set_backend_module
package main

import (
	"C"
	"bufio"
	"log"
	"net"
	"strings"
)

const (
	socket = "/lab/build/ngx_http_set_backend.sock"
)

//export LookupBackend
func LookupBackend(_host *C.char) *C.char {
	host := C.GoString(_host)

	c, err := net.Dial("unix", socket)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	//ask for the backend
	_, err = c.Write([]byte(host + "\n"))
	if err != nil {
		log.Fatal(err)
	}

	// read the answer
	backend, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	backend = strings.TrimSuffix(backend, "\n")

	return C.CString(string(backend))
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
