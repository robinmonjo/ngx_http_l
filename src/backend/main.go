package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	socket  = "/lab/build/ngx_http_set_backend.sock"
	dbFile  = "/lab/build/backends.db"
	nobody  = "nobody"
	apiHost = "l.io"
	apiPort = "9999"
)

func main() {
	provider := &provider{
		socket:          socket,
		username:        nobody,
		dbFile:          dbFile,
		internalMapping: map[string]string{apiHost: "127.0.0.1:9999"},
	}
	defer provider.cleanup()
	go func() {
		if err := provider.listen(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := startApi(apiPort); err != nil {
			log.Fatal(err)
		}
	}()

	// signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
