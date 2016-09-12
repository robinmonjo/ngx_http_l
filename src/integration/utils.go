package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func apiDo(verb, endpoint string, payload map[string]string) (interface{}, int, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(payload)

	req, err := http.NewRequest(verb, endpoint, b)
	req.Host = fmt.Sprintf("api.%s", os.Getenv("DOMAIN"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var res interface{}
	if err := decoder.Decode(&res); err != nil {
		return nil, 0, err
	}

	return res, resp.StatusCode, nil
}
