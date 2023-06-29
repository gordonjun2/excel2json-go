package reader

import (
	"fmt"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/gordonjun2/excel2json-go/pkg/utils"
	"github.com/gordonjun2/excel2json-go/pkg/parser"
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
	if byteFile, err = utils.GetFilePath(path); err != nil {
		return nil, err
	}
	if result, err = parser.ParseExcelFileData(path, byteFile, sheetName, localCache); err != nil {
		return nil, err
	}
	if len(headers) > 0 {
		result = utils.FilterDataHeaders(utils.MapHeaders(headers), result)
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
		keyName  = utils.HashKeyString(fmt.Sprintf(`%s||%s||%s`, path, delimiter, strings.Join(headers, "||")))
	)
	if byteFile, err = utils.GetFilePath(path); err != nil {
		return nil, err
	}
	if result, err = parser.ParseCsvFileData(byteFile, delimiter, keyName, localCache); err != nil {
		return nil, err
	}
	if len(headers) > 0 {
		result = utils.FilterDataHeaders(utils.MapHeaders(headers), result)
	}
	return result, nil
}
