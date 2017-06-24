package main

import (
	"flag"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	u "github.com/hiromaily/golibs/utils"
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
  gogentype -json '{"str": "xxxx", "slice": [1,2,3], "sliceempty": [], "null": null, "int": 10, "zero": 0, "bool": true, "obj": {"child":100}}'

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
	fmt.Printf("type %s struct {\n", "TypeName")
	for key, value := range unmarshaledJson {
		//lg.Debug("key:", key, " value:", value)
		name := strings.Title(key)
		typeStr := u.CheckInterfaceByIf(value)
		if typeStr == "" || typeStr == "slice" {
			typeStr = "*string"
		} else if typeStr == "float64" {
			typeStr = "int"
		}
		fmt.Printf("\t%s\t%s\t`json:\"%s\"`\n", name, typeStr, key)
	}
	fmt.Println("}")
}
