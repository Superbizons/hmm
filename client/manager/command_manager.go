package manager

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

func SendCommand(cmd interface{}, url string) (*os.File, error) {
	js, err := json.Marshal(cmd)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(js)))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}

	filename := resp.Header.Get("X-FileName")

	if strings.Contains(filename, "/") {
		return nil, errors.New("Invalid file name." + filename)
	}

	botfile, err := os.Create(filename)

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(botfile, resp.Body)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return botfile, nil
}
