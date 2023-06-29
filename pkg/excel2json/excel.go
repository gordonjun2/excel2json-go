package excel2json

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// parseExcelFileData to read all data excel
// data is byte data from http client
// keyName is string concat
// sheetName is string
func parseExcelFileData(filePath string, data []byte, sheetName string) ([]*map[string]interface{}, error) {
	var (
		headers []string
		result  []*map[string]interface{}
		wb      = new(excelize.File)
		err     error
		keyName string
	)

	if sheetName != "" {
		keyName = hashKeyString(fmt.Sprintf(`%s||%s||%s`, filePath, sheetName, strings.Join(headers, "||")))

		// if already data in cache
		if cacheData, found := localCache.Get(keyName); found {
			return cacheData.([]*map[string]interface{}), nil
		}
		// open byte data with reader
		if wb, err = excelize.OpenReader(bytes.NewReader(data)); err != nil {
			return nil, err
		}

	} else {
		// open byte data with reader
		if wb, err = excelize.OpenReader(bytes.NewReader(data)); err != nil {
			return nil, err
		}

		sheetName = wb.GetSheetName(1)

		keyName = hashKeyString(fmt.Sprintf(`%s||%s||%s`, filePath, sheetName, strings.Join(headers, "||")))

		// if already data in cache
		if cacheData, found := localCache.Get(keyName); found {
			return cacheData.([]*map[string]interface{}), nil
		}
	}

	// Get all the rows in the Sheet.
	rows := wb.GetRows(sheetName)
	headers = rows[0]
	for _, row := range rows[1:] {
		var tmpMap = make(map[string]interface{})
		for j, v := range row {
			tmpMap[strings.Join(strings.Split(headers[j], " "), "")] = v
		}
		result = append(result, &tmpMap)
	}
	// set data to cache
	localCache.Set(keyName, result, 10*time.Minute)
	return result, nil
}
