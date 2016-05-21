package integration

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

// helpers
const (
	port = "8888"
)

func TestMain(m *testing.M) {
	backend := &backend{}
	go func() {
		if err := backend.start(); err != nil {
			log.Fatalf("unable to start backend %v", err)
		}
	}()
	time.Sleep(1 * time.Second)
	if err := startNginx(); err != nil {
		log.Fatalf("unable to start nginx %v", err)
	}

	exitCode := m.Run()

	if err := backend.stop(); err != nil {
		log.Fatalf("unable to stop backend %v", err)
	}
	if err := stopNginx(); err != nil {
		log.Fatalf("unable to stop nginx %v", err)
	}

	os.Exit(exitCode)
}

// integration testing
func TestNginxRunning(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s", port))
	if err != nil {
		t.Fatal(err)
	}
	code := 200
	if resp.StatusCode != code {
		t.Fatalf("expected status code to be %d, got %d", code, resp.StatusCode)
	}

	fmt.Println("All good baby")
}
