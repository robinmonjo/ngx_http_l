package main

import (
	"bufio"
	"net"
	"os"
	"strings"

	"github.com/boltdb/bolt"
)

type provider struct {
	socket          string
	username        string
	listener        net.Listener
	db              *bolt.DB
	internalMapping map[string]string
}

func (p *provider) listen() error {
	var err error

	//create unix socket
	os.RemoveAll(p.socket) //in case it was not destroyed properly on exit
	p.listener, err = net.Listen("unix", p.socket)
	if err != nil {
		return err
	}
	if err = chown(p.username, p.socket); err != nil {
		return err
	}

	//start listening
	for {
		conn, err := p.listener.Accept()
		if err != nil {
			return err
		}
		go func(c net.Conn) {
			defer c.Close()
			if err := p.processRequest(c); err != nil {
				//TODO handle error
			}
		}(conn)
	}
}

func (p *provider) processRequest(c net.Conn) error {
	host, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		return err
	}
	host = strings.TrimSuffix(host, "\n")

	backend := p.internalMapping[host]
	if backend == "" {
		backend, err = p.lookupBackend(host)
		if err != nil {
			return err
		}
	}

	_, err = c.Write([]byte(backend + "\n"))
	return err
}

func (p *provider) lookupBackend(host string) (string, error) {
	backend := "www.google.com"
	if err := p.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}
		res := b.Get([]byte(host))
		if res != nil {
			backend = string(res)
		}
		return nil
	}); err != nil {
		return "", err
	}

	return backend, nil
}

func (p *provider) cleanup() {
	p.listener.Close()
	os.RemoveAll(p.socket)
	p.db.Close()
}
