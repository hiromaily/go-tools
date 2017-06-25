package main

import (
	"flag"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	u "github.com/hiromaily/golibs/utils"
	tm "github.com/hiromaily/golibs/time"
	//"io/ioutil"
	"encoding/json"
	"os"
	"strings"
)

var (
	jsonString = flag.String("json", "", "Json String Data")
)

var usage = `Usage: %s [options...]
Options:
  -json  Package name.
e.g.:
  gogentype -json '{"str": "xxxx", "slice": [1,2,3], "sliceempty": [], "null": null, "int": 10, "zero": 0, "bool": true, "date": "2017-07-26T11:10:15+02:00", "obj": {"child":100}}'

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

	if *jsonString == "" {
		flag.Usage()

		os.Exit(1)
		return
	}
	//lg.Debug(*jsonString)
}

func main() {
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
			typeStr = "[]string"
		} else if typeStr == "float64" {
			typeStr = "int"
		} else if typeStr == "map" {
			//change interface{} to []map[string]interface{}
			mi := u.ItoMsif(value)
			if mi != nil{
				jsonMaps = append(jsonMaps, mi)
				typeStr = fmt.Sprintf("TypeName%d", idx+1)
			}
		}
		fmt.Printf("\t%s\t%s\t`json:\"%s\"`\n", name, typeStr, key)
	}
	fmt.Println("}")

	//recursion
	for i, v := range jsonMaps {
		printType(v, idx+i+1)
	}
}

func isDateTime(){
	//s := "2015/12/22 10:00:30"
	//layout := "2006/01/02 15:04:05"
	//t, err := time.Parse(layout, s)
	//if err != nil {
	//	panic(err)
	//}
}

//func changeMapToJson(data map[string]interface{}) ([]byte, error){
//	fmt.Println("data is", data)
//	b, err := json.Marshal(data)
//	if err != nil {
//		return nil, fmt.Errorf("[ERROR] When calling `json.Marshal`: %v\n", err)
//	}
//
//	return b, nil
//}