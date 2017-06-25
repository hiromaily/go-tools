package main

import (
	"flag"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	u "github.com/hiromaily/golibs/utils"
	tm "github.com/hiromaily/golibs/time"
	"encoding/json"
	"os"
	"strings"
	"io/ioutil"
	"errors"
)

var (
	jsonString = flag.String("json", "", "Json String Data")
	file       = flag.String("file", "", "Json File Path")
)

var usage = `Usage: %s [options...]
Options:
  -json  Json String Data.
e.g.:
  gogentype -json '{"str": "xxxx", "slice": [1,2,3], "sliceempty": [], "null": null, "int": 10, "zero": 0, "bool": true, "date": "2017-07-26T11:10:15+02:00", "obj": {"child":100}}'
 or
  gogentype -file sample.json

Note:null value can not be detected proper type.
`

// Params is parameter for template file
type Params struct {
	Name      string
	Uppercase string
}

func init() {
	lg.InitializeLog(lg.DebugStatus, lg.LogOff, 99, "[GOTOOLS GoGenType]", "/var/log/go/gotool.log")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, os.Args[0]))
	}

	flag.Parse()

	if *jsonString == "" && *file == "" {
		flag.Usage()

		os.Exit(1)
		return
	}
	//lg.Debug(*jsonString)
}

func main() {
	if *jsonString != ""{
		handleJsonData()
	}
	if *file != ""{
		handleJsonFile()
	}
}

func handleJsonData(){
	var unmarshaledJson map[string]interface{}

	//1. json
	err := json.Unmarshal([]byte(*jsonString), &unmarshaledJson)
	if err != nil {
		lg.Errorf("After calling json.Unmarshal(): %v", err)
		return
	}
	//lg.Debug(unmarshaledJson)

	//2. handle response and output json
	printType(unmarshaledJson, 1)
}

func handleJsonFile(){
	var unmarshaledJson map[string]interface{}

	//1. json load
	jsonByte, err := loadJSONFile(*file)
	if err != nil {
		lg.Errorf("After calling loadJSONFile(): %v", err)
		return
	}

	//2. json
	err = json.Unmarshal(jsonByte, &unmarshaledJson)
	if err != nil {
		lg.Errorf("After calling json.Unmarshal(): %v", err)
		return
	}
	//lg.Debug(unmarshaledJson)

	//3. handle response and output json
	printType(unmarshaledJson, 1)

}


func printType(jsonData map[string]interface{}, idx int){
	var jsonMaps []map[string]interface{}

	fmt.Printf("type %s%d struct {\n", "TypeName", idx)
	for key, value := range jsonData {
		//lg.Debug("key:", key, " value:", value)
		name := strings.Title(key)
		typeStr := u.CheckInterfaceByIf(value)
		if typeStr == "" {
			typeStr = "*string"
		} else if typeStr == "string"{
			//Check if it's date. e.g. 2017-07-26T11:10:15+02:00
			ret := tm.CheckParseTime(u.Itos(value))
			if len(ret) > 0{
				typeStr = "*time.Time"
			}
		} else if typeStr == "slice"{
			//check value
			typeStr, value = handleSliceValue(value)
			if typeStr == "[]map"{
				typeStr, jsonMaps = handleMapValue(value, jsonMaps, idx)
			}
		} else if typeStr == "float64" {
			typeStr = "int"
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

func handleSliceValue(value interface{}) (string, interface{}){
	sliceIfc := u.ItoSI(value)
	if sliceIfc != nil && len(sliceIfc) > 0{
		//check type
		typeStr := u.CheckInterfaceByIf(sliceIfc[0])
		return fmt.Sprintf("[]%s", typeStr), sliceIfc[0]
	}
	//In nil value, return string for now.
	return "[]string", nil
}

func handleMapValue(value interface{}, jsonMaps []map[string]interface{}, idx int) (string, []map[string]interface{}){
	var typeStr string
	//change interface{} to []map[string]interface{}
	mi := u.ItoMsif(value)
	if mi != nil{
		jsonMaps = append(jsonMaps, mi)
		typeStr = fmt.Sprintf("TypeName%d", idx+1)
	}else{
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