package manager

import (
	"encoding/json"
	"io/ioutil"

	"../basic"
)

var (
	Configuration *basic.Configuration
)

func LoadConfiguration() (*basic.Configuration, error) {
	configurationFile, err := ioutil.ReadFile("./configuration.json")

	if err != nil {
		return &basic.Configuration{}, err
	}

	err = json.Unmarshal(configurationFile, &Configuration)

	if err != nil {
		return Configuration, err
	}

	return Configuration, nil
}
