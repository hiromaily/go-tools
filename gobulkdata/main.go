package main

import (
	"bufio"
	"flag"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	"os"
	"strconv"
	"strings"
)

type TestData struct {
	id          int
	q1          string
	q2          string
	q3          string
	q4          string
	q5          string
	q6          string
	q7          string
	q8          string
	q9          string
	q10         string
	q11         string
	startedAt   string
	completedAt string
	platform    string
	deviceType  string
	deviceBrand string
	browser     string
	ip          string
	email       string
	hiddenEmail string
	hiddenName  string
}

var (
	fileName = flag.String("f", "", "Filename")
	lineNum  = flag.Int("l", 10, "Max Line number")
	//testData = []TestData{}
	testData = TestData{}
	//initID     = 2500
	//mailFormat = "gogo%d@gmail.com"
	//maxLines   = 50001
)

var usage = `Usage: %s [options...]
Options:
  -f  File name.
  -t  File type.
e.g.:
  gobulkdata -f ${HOME}/work/go/src/github.com/hiromaily/gotools/text.txt -l 20
`

//TODO
//add file type
//file can not be output

//text.txt
//2501,,,,,,,,,,Newsletter,Artist c,2016-06-15 13:18:43,2016-06-15 13:19:13,Windows,Desktop,unknown,Chrome,92.111.79.210,gogo1@gmail.com,,test name
//2502,,,,,,,,,,Newsletter,Artist c,2016-06-15 13:18:43,2016-06-15 13:19:13,Windows,Desktop,unknown,Chrome,92.111.79.210,gogo2@gmail.com,,test name
//2503,,,,,,,,,,Newsletter,Artist c,2016-06-15 13:18:43,2016-06-15 13:19:13,Windows,Desktop,unknown,Chrome,92.111.79.210,gogo3@gmail.com,,test name

func init() {
	lg.InitializeLog(lg.DebugStatus, lg.LogOff, 99, "[GOTOOLS GOTEST]", "/var/log/go/gotool.log")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, os.Args[0]))
	}

	flag.Parse()

	if *fileName == "" {
		flag.Usage()

		os.Exit(1)
		return
	}

	//initial value
	//testData = append(testData, TestData{
	testData = TestData{
		id:          2500,
		q1:          "",
		q2:          "",
		q3:          "",
		q4:          "",
		q5:          "",
		q6:          "",
		q7:          "",
		q8:          "",
		q9:          "",
		q10:         "Newsletter",
		q11:         "Artist c",
		startedAt:   "2016-06-15 13:18:43",
		completedAt: "2016-06-15 13:19:13",
		platform:    "Windows",
		deviceType:  "Desktop",
		deviceBrand: "unknown",
		browser:     "Chrome",
		ip:          "92.111.79.210",
		email:       "gogo%d@gmail.com",
		hiddenEmail: "",
		hiddenName:  "test name",
	}
}

func makeData() {
	var writer *bufio.Writer

	fmt.Println(*fileName)

	file, _ := os.OpenFile(*fileName, os.O_WRONLY, 0644)
	writer = bufio.NewWriter(file)

	for i := 1; i < *lineNum; i++ {
		tmpData := []string{
			strconv.Itoa(testData.id + i),
			testData.q1,
			testData.q2,
			testData.q3,
			testData.q4,
			testData.q5,
			testData.q6,
			testData.q7,
			testData.q8,
			testData.q9,
			testData.q10,
			testData.q11,
			testData.startedAt,
			testData.completedAt,
			testData.platform,
			testData.deviceType,
			testData.deviceBrand,
			testData.browser,
			testData.ip,
			fmt.Sprintf(testData.email, i),
			testData.hiddenEmail,
			testData.hiddenName,
		}
		fmt.Println(strings.Join(tmpData[:], ","))
		writer.WriteString(strings.Join(tmpData[:], ",") + "\n")
	}

	writer.Flush()

}

func main() {
	makeData()
}
