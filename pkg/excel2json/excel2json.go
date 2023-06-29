package excel2json

import (
	"fmt"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

// Create a cache with a default expiration time of 5 minutes, and which
// purges expired items every 10 minutes
var localCache = cache.New(5*time.Minute, 10*time.Minute)

// GetExcelFilePath to read all data excel
// path is string path file name
// sheetName is sheet name in excel file
// Headers is array string that is want to select specific data
func GetExcelFilePath(path, sheetName string, headers []string) ([]*map[string]interface{}, error) {
	var (
		result   []*map[string]interface{}
		err      error
		byteFile []byte
	)
	if byteFile, err = getFilePath(path); err != nil {
		return nil, err
	}
	if result, err = parseExcelFileData(path, byteFile, sheetName); err != nil {
		return nil, err
	}
	if len(headers) > 0 {
		result = filterDataHeaders(mapHeaders(headers), result)
	}
	return result, nil
}

// GetCsvFilePath to read all data excel
// path is string path file name
// sheetName is sheet name in excel file
// Headers is array string that is want to select specific data
func GetCsvFilePath(path, delimiter string, headers []string) ([]*map[string]interface{}, error) {
	var (
		result   []*map[string]interface{}
		err      error
		byteFile []byte
		keyName  = hashKeyString(fmt.Sprintf(`%s||%s||%s`, path, delimiter, strings.Join(headers, "||")))
	)
	if byteFile, err = getFilePath(path); err != nil {
		return nil, err
	}
	if result, err = parseCsvFileData(byteFile, delimiter, keyName); err != nil {
		return nil, err
	}
	if len(headers) > 0 {
		result = filterDataHeaders(mapHeaders(headers), result)
	}
	return result, nil
}
