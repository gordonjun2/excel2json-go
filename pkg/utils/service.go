package utils

import (
	"io/ioutil"
)

// GetFilePath fetch file body with http client
// filePath string url
func GetFilePath(filePath string) ([]byte, error) {
	var (
		body []byte
		err  error
	)
	// just pass the file name
	if body, err = ioutil.ReadFile(filePath); err != nil {
		return nil, err
	}
	return body, nil
}
