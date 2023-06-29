package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

// MapHeaders Convert all headers to map string
// headers is array string
func MapHeaders(headers []string) map[string]string {
	var result = make(map[string]string)
	for _, val := range headers {
		result[strings.Join(strings.Split(val, " "), "")] = strings.Join(strings.Split(val, " "), "")
	}
	return result
}

// FilterDataHeaders data based on selected header
// mapHeader is headers that already mapped
// data is data that is already fetch from excel file
func FilterDataHeaders(mapHeader map[string]string, data []*map[string]interface{}) []*map[string]interface{} {
	var result []*map[string]interface{}
	for _, val := range data {
		var (
			t = make(map[string]interface{})
			d = *val
		)
		for _, header := range mapHeader {
			if _, ok := d[header]; ok {
				t[header] = d[header]
			}
		}
		result = append(result, &t)
	}
	return result
}

// HashKeyString is method to hash keyName
// keyName is string
func HashKeyString(keyName string) string {
	var hashString = sha1.New()
	hashString.Write([]byte(keyName))
	return base64.URLEncoding.EncodeToString(hashString.Sum(nil))
}
