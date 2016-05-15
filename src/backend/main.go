package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	socket = "/lab/build/ngx_http_set_backend.sock"
	dbFile = "/lab/build/backends.db"
	nobody = "nobody"
)

func main() {
	provider := &provider{
		socket:   socket,
		username: nobody,
		dbFile:   dbFile,
	}
	defer provider.cleanup()
	go func() {
		if err := provider.listen(); err != nil {
			log.Fatal(err)
		}
	}()

	// signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
