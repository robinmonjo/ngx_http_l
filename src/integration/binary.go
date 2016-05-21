package integration

import (
	"os"
	"os/exec"
	"syscall"
)

const (
	nginx      = "/lab/build/nginx/sbin/nginx"
	backendBin = "/lab/build/backend"
)

// nginx
func startNginx() error {
	return exec.Command(nginx).Run()
}

func stopNginx() error {
	return exec.Command(nginx, "-s", "stop").Run()
}

// backend
type backend struct {
	ps *os.Process
}

func (b *backend) start() error {
	cmd := exec.Command(backendBin)
	if err := cmd.Start(); err != nil {
		return err
	}
	b.ps = cmd.Process
	return cmd.Wait()
}

func (b *backend) stop() error {
	return b.ps.Signal(syscall.SIGTERM)
}
