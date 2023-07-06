package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gordonjun2/excel2json-go/pkg/reader"
)

func main() {
	var (
		result []*map[string]interface{}
		err    error
		path   string
		// select only selected field
		// if you want to show all headers just passing nil or empty list
		// eg. headers   = []string{"humidity", "sound"}
		headers    = []string{}
		delimited  = ","
		jsonString = "["
	)

	if len(os.Args) >= 2 {
		path = os.Args[1]
	}

	if path == "" {
		log.Fatalf("No file path provided!")
	} else {
		fmt.Println("Input file:")
		fmt.Println(path)
	}

	if result, err = reader.GetCsvFilePath(path, delimited, headers); err != nil {
		log.Fatalf(`unable to parse file, error: %s`, err)
	}

	fmt.Println("\nOutput JSON:")

	for i, val := range result {
		result, _ := json.Marshal(val)
		if i != 0 {
			jsonString = jsonString + ", "
		}
		jsonString = jsonString + string(result)
	}

	jsonString = jsonString + "]"

	fmt.Println(jsonString)
}
