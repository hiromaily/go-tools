package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	lg "github.com/hiromaily/golibs/log"
	"github.com/hiromaily/golibs/tmpl"
)

var (
	name = flag.String("n", "", "package name")
)

var usage = `Usage: %s [options...]
Options:
  -n  package name.
e.g.:
  gen-testfile -n new_pkg_name
   >> new_pgk_name_test.go
`

// Params is parameter for template file
type Params struct {
	Name      string
	Uppercase string
}

func init() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, os.Args[0]))
	}
	flag.Parse()

	if *name == "" {
		flag.Usage()
		os.Exit(1)
		return
	}
}

//change uppercase only first character
func uppercase(str string) string {
	result := ""
	for i, bs := range str {
		s := string(bs)
		if i == 0 {
			result = strings.ToUpper(s)
			continue
		}
		result += s
	}
	return result
}

func readTemplate() string {
	goPath := os.Getenv("GOPATH")
	tpl, err := template.ParseFiles(goPath + "/src/github.com/hiromaily/gotools/go-testfile/templates/base.tpl")
	if err != nil {
		lg.Fatalf("parse error 1: %s", "templates/base.tpl")
	}

	params := Params{Name: *name, Uppercase: uppercase(*name)}

	result, err := tmpl.FileTempParser(tpl, params)
	if err != nil {
		lg.Fatalf("parse error 2: %s", "templates/base.tpl")
	}

	return result
}

func outFile(data string) {
	ioutil.WriteFile(fmt.Sprintf("./%s_test.go", *name), []byte(data), 0644)
}

func main() {
	lg.InitializeLog(lg.DebugStatus, lg.TimeShortFile, "[GOTOOLS GoTestFile]", "", "hiromaily")

	result := readTemplate()

	//output as file
	outFile(result)

	//lg.Debug(result)
}
