package common

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"errors"
)

// check string is empty
func IsEmptyString(s string) bool {
	return s == ""
}

// WriteDataToFile write data to file
func WriteDataToFile(filename string, data []byte) error {
	if IsEmptyString(filename) {
		return errors.New("filename invalid input")
	}
	path := filepath.Dir(filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0600)
	}

	return ioutil.WriteFile(filename, data, 0600)
}
