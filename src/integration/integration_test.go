package integration

import (
	"fmt"
	"net/http"
	"os/exec"
	"testing"
)

// helpers
const (
	nginx    = "/lab/build/nginx/sbin/nginx"
	backends = "/lab/build/backend"
	port     = "8888"
)

func startNginx(t *testing.T) {
	if err := exec.Command(nginx).Run(); err != nil {
		t.Fatal(err)
	}
	if err := exec.Command(backends).Run(); err != nil {
		t.Fatal(err)
	}
}

func stopNginx(t *testing.T) {
	if err := exec.Command(nginx, "-s", "stop").Run(); err != nil {
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
