package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	lg "github.com/hiromaily/golibs/log"
	tm "github.com/hiromaily/golibs/time"
	u "github.com/hiromaily/golibs/utils"
)

var (
	jsonString = flag.String("json", "", "Json String Data")
	file       = flag.String("file", "", "Json File Path")
)

var usage = `Usage: %s [options...]
Options:
  -json  Json String Data.
e.g.:
  gen-struct -json '{"str": "xxxx", "slice": [1,2,3], "sliceempty": [], "null": null, "int": 10, "zero": 0, "bool": true, "date": "2017-07-26T11:10:15+02:00", "obj": {"child":100}}'
 or
  gen-struct -file sample.json

Note:null value can not be detected proper type.
`

func init() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, os.Args[0]))
	}
	flag.Parse()

	if *jsonString == "" && *file == "" {
		flag.Usage()

		os.Exit(1)
		return
	}
}

func main() {
	lg.InitializeLog(lg.DebugStatus, lg.TimeShortFile, "[GOTOOLS GoGenType]", "", "hiromaily")

	if *jsonString != "" {
		handleJSONData()
	}
	if *file != "" {
		handleJSONFile()
	}
}

func handleJSONData() {
	var unmarshaledJSON map[string]interface{}

	//1. json
	err := json.Unmarshal([]byte(*jsonString), &unmarshaledJSON)
	if err != nil {
		lg.Errorf("After calling json.Unmarshal(): %v", err)
		return
	}
	//lg.Debug(unmarshaledJson)

	//2. handle response and output json
	printType(unmarshaledJSON, 1)
}

func handleJSONFile() {
	var unmarshaledJSON map[string]interface{}

	//1. json load
	jsonByte, err := loadJSONFile(*file)
	if err != nil {
		lg.Errorf("After calling loadJSONFile(): %v", err)
		return
	}

	//2. json
	err = json.Unmarshal(jsonByte, &unmarshaledJSON)
	if err != nil {
		lg.Errorf("After calling json.Unmarshal(): %v", err)
		return
	}

	//3. handle response and output json
	printType(unmarshaledJSON, 1)

}

func printType(jsonData map[string]interface{}, idx int) {
	var jsonMaps []map[string]interface{}

	fmt.Printf("type %s%d struct {\n", "TypeName", idx)
	for key, value := range jsonData {
		//lg.Debug("key:", key, " value:", value)
		name := strings.Title(key)

		typeStr := u.CheckInterfaceByIf(value)
		typeStr = typeConvert(typeStr, value)

		if typeStr == "slice" {
			//check value
			typeStr, value = handleSliceValue(value)
			if typeStr == "[]map" {
				typeStr, jsonMaps = handleMapValue(value, jsonMaps, idx)
			}
		} else if typeStr == "map" {
			typeStr, jsonMaps = handleMapValue(value, jsonMaps, idx)
		}

		fmt.Printf("\t%s\t%s\t`json:\"%s\"`\n", name, typeStr, key)
	}
	fmt.Println("}")

	//recursion
	for i, v := range jsonMaps {
		printType(v, idx+i+1)
	}
}

func typeConvert(typeStr string, value interface{}) string {
	if typeStr == "" {
		typeStr = "*string"
	} else if typeStr == "string" {
		//Check if it's date. e.g. 2017-07-26T11:10:15+02:00
		ret := tm.CheckParseTime(u.Itos(value))
		if len(ret) > 0 {
			typeStr = "*time.Time"
		}
	} else if typeStr == "float64" {
		typeStr = "int"
	}
	return typeStr
}

func handleSliceValue(value interface{}) (string, interface{}) {
	sliceIfc := u.ItoSI(value)
	if len(sliceIfc) > 0 {
		//check type
		typeStr := u.CheckInterfaceByIf(sliceIfc[0])
		typeStr = typeConvert(typeStr, sliceIfc[0])

		return fmt.Sprintf("[]%s", typeStr), sliceIfc[0]
	}
	//In nil value, return string for now.
	return "[]string", nil
}

func handleMapValue(value interface{}, jsonMaps []map[string]interface{}, idx int) (string, []map[string]interface{}) {
	var typeStr string
	//change interface{} to []map[string]interface{}
	mi := u.ItoMsif(value)
	if mi != nil {
		jsonMaps = append(jsonMaps, mi)
		typeStr = fmt.Sprintf("TypeName%d", idx+1)
	} else {
		typeStr = "map"
	}
	return typeStr, jsonMaps
}

// LoadJSONFile is to read json file
func loadJSONFile(filePath string) ([]byte, error) {
	// Loading jsonfile
	if filePath == "" {
		err := errors.New("nothing JSON file")
		return nil, err
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
