package common

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"errors"
	"regexp"
	"fmt"
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

func ParseData(regex, data string) (string, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return "", fmt.Errorf("failed to create regex expresion, regex: %s, err: %v", regex, err)
	}
	result := re.FindStringSubmatch(data)
	if len(result) < 2 {
		return "", fmt.Errorf("unable to find sub string")
	}
	return result[1], nil
}
