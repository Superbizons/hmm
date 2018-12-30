package manager

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func SendCommand(cmd interface{}, url string) error {
	js, err := json.Marshal(cmd)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(js)))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.Status != "200 OK" {
		return errors.New(resp.Status)
	}

	defer resp.Body.Close()
	return nil
}
