package script

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// script
type Script interface {
	Exec(command string) ([]byte, error)
}

func createScriptFile(in,out string) (string, error) {

	data, err := Asset(in)
	if err != nil {
		return "", err
	}

	dir := os.TempDir()
	path := filepath.Join(dir, out)
	if exists(path) {
		return path, nil
	}

	err = ioutil.WriteFile(path, data, 777)
	if err != nil {
		return "", err
	}
	return path, nil
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

