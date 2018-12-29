package manager

import "os"

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
