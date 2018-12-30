package manager

import (
	"io/ioutil"
	"os"
	"strings"
)

func CreateFileIfnExist(file string) error {
	_, err := os.Stat(file)

	if os.IsNotExist(err) {
		_, err = os.Create(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateDirIfnExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)

		if err != nil {
			return err
		}
	}
	return nil
}

func GetBotZipsFromDir(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)

	var zipfiles = make([]string, 0)

	if err != nil {
		return zipfiles, err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".zip") {
			continue
		}

		zipfiles = append(zipfiles, file.Name())
	}

	return zipfiles, nil
}
