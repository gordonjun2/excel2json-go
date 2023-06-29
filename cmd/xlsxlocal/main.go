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
		result    []*map[string]interface{}
		err       error
		path      string
		sheetName string
		// select only selected field
		// if you want to show all headers just passing nil or empty list
		// eg. headers = []string{"Profit", "Shipping Cost", "Unit Price"}
		headers  = []string{}
		jsonList = []string{}
	)

	if len(os.Args) >= 3 {
		path = os.Args[1]
		sheetName = os.Args[2]
	} else if len(os.Args) == 2 {
		path = os.Args[1]
	}

	if path == "" {
		log.Fatalf("No file path provided!")
	} else {
		fmt.Println("Input file:")
		fmt.Println(path)
	}

	if result, err = reader.GetExcelFilePath(path, sheetName, headers); err != nil {
		log.Fatalf(`unable to parse file, error: %s`, err)
	}

	fmt.Println("\nOutput JSON:")

	for _, val := range result {
		result, _ := json.Marshal(val)
		jsonList = append(jsonList, string(result))
	}

	fmt.Println(jsonList)
}
