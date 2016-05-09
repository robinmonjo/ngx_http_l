package integration

import (
	"fmt"
	"net/http"
	"os/exec"
	"testing"
)

// helpers
const (
	binary = "/lab/build/nginx/sbin/nginx"
	port   = "8888"
)

func startNginx(t *testing.T) {
	fmt.Printf("starting nginx ...")
	if err := exec.Command(binary).Run(); err != nil {
		t.Fatal(err)
	}
	fmt.Println("done")
}

func stopNginx(t *testing.T) {
	if err := exec.Command(binary, "-s", "stop").Run(); err != nil {
		t.Fatal(err)
	}
}

// integration testing
func TestNginxRunning(t *testing.T) {
	startNginx(t)
	defer stopNginx(t)

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
