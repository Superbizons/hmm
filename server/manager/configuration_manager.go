package manager

import (
	"encoding/json"
	"io/ioutil"

	"../api"
)

var (
	Configuration *api.Configuration
)

func LoadConfiguration() (*api.Configuration, error) {
	configurationFile, err := ioutil.ReadFile("./configuration.json")

	if err != nil {
		return &api.Configuration{}, err
	}

	err = json.Unmarshal(configurationFile, &Configuration)

	if err != nil {
		return Configuration, err
	}

	return Configuration, nil
}
