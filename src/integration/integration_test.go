package integration

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

// helpers
const (
	nginxURL = "http://127.0.0.1:8888"
)

func TestMain(m *testing.M) {
	backend := &backend{}
	go func() {
		if err := backend.start(); err != nil {
			log.Fatalf("unable to start backend %v", err)
		}
	}()
	time.Sleep(100 * time.Millisecond) //ensure backend started
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
func TestSimpleGet(t *testing.T) {
	resp, err := http.Get(nginxURL)
	if err != nil {
		t.Fatal(err)
	}
	code := 200
	if resp.StatusCode != code {
		t.Fatalf("expected status code to be %d, got %d", code, resp.StatusCode)
	}
}

// rest API testing
func TestCreate(t *testing.T) {
	const (
		host    = "host1.com"
		backend = "10.0.3.10"
	)

	var payload = map[string]string{
		"host":    host,
		"backend": backend,
	}

	_, code, err := apiDo("POST", fmt.Sprintf("%s/entries.json", nginxURL), payload)
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := 200
	if code != expectedCode {
		t.Fatalf("expected status code to be %d, got %d", expectedCode, code)
	}
}

func TestIndex(t *testing.T) {
	body, code, err := apiDo("GET", fmt.Sprintf("%s/entries.json", nginxURL), nil)
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := 200
	if code != expectedCode {
		t.Fatalf("expected status code to be %d, got %d", expectedCode, code)
	}
	entries := body.([]interface{})
	if len(entries) != 1 {
		t.Fatalf("expected 1 entries, got %d", len(entries))
	}
}

func TestDestroy(t *testing.T) {
	host := url.QueryEscape("host1.com")
	body, code, err := apiDo("DELETE", fmt.Sprintf("%s/entries/%s.json", nginxURL, host), nil)
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := 200
	if code != expectedCode {
		t.Fatalf("expected status code to be %d, got %d", expectedCode, code)
	}
}

// provider
func TestProvider(t *testing.T) {
	host := "success.com"
	backend := "www.apple.com"

	var payload = map[string]string{
		"host":    host,
		"backend": backend,
	}

	_, _, err := apiDo("POST", fmt.Sprintf("%s/entries.json", nginxURL), payload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/library/test/success.html", nginxURL), nil)
	req.Host = host
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "<HTML><HEAD><TITLE>Success</TITLE></HEAD><BODY>Success</BODY></HTML>" {
		t.Fatalf("not the expected content, got: %q", string(content))
	}
}
