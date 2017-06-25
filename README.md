# gotools

[![Go Report Card](https://goreportcard.com/badge/github.com/hiromaily/gotools)](https://goreportcard.com/report/github.com/hiromaily/gotools)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/hiromaily/gotools/master/LICENSE)

go tools


# 1. gocipher [For Development]
Create encrypted string both (encode and decode)

## Installation
```
$ go get github.com/hiromaily/gotools/gocipher
```

It requires below environment variable, e.g. 
```
export 'ENC_KEY=M#XF#R+gaKFvJ_<'
export 'ENC_IV=@~wK-3OlQ<c2y@DA'
```

## Usage
```
Options:
  -m  e:encode, d:decode.
e.g.:
  gocipher -m e xxxxxxxx
    or
  gocipher -m d xxxxxxxx
```

## Example
```
$ gocipher -m e secret_string
 => gtBl3kNqSAJGvJjnvUU9HQ==

$ gocipher -m d gtBl3kNqSAJGvJjnvUU9HQ==
 => secret_string
```


# 2. gocookie [For Development]
Retrieve cookie data by domain from chrome.  
(This is not my development.)

## Usage
```
e.g.:
  gocookie domain.com
```

## Example
```
$ gocookie localhost
 => localhost/key: value
```


# 3. goapitest [For Development]
This is for API Test

## Usage
```
e.g.:
  goapitest -m user
```

## Example
```
$ goapitest -m user
 => 
{
  "success": true,
  "records": [
    {
      "id": "xxxxxx",
      "attr": null,
      "name": "firstlast0013",
    }
  ]
}
```

## Create New API Endpoint
1. add model file in models/  
   It's easier to define structure by `gogentype` in gotools with json raw data.
2. add condition at main() in main.go. 


# 4. gotestfile [For Golang Development]
Create template file of xxx_test.go

## Installation
```
$ go get github.com/hiromaily/gotools/gotestfile
```

## Usage
```
Usage: gotestfile [options...]

Options:
  -n     package name

e.g.
 $ gotestfile -n new-package-name
```

## Example
```
$ gotestfile -n newpkg
 => generate newpkg_test.go file
```
```
package newpkg_test

import (
	. "github.com/hiromaily/golibs/newpkg"
	//lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	"os"
	"testing"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Newpkg]")
}

func setup() {
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// function
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestNewpkg(t *testing.T) {
	//if err != nil {
	//	t.Errorf("TestNewpkg error: %s", err)
	//}
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkNewpkg(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
```


# 5. gogentype [For Golang Development]
Create golang type struct from json data.

## Installation
```
$ go get github.com/hiromaily/gotools/gogentype
```

## Usage
```
Usage: gogentype [options...]
Options:
  -json  Package name.
e.g.:
  gogentype -json '{"str": "xxxx", "slice": [1,2,3], "sliceempty": [], "null": null, "int": 10, "zero": 0, "bool": true, "obj": {"child":100}}'
 or
  gogentype -file sample.json

Note:null value can not be detected proper type.
```

## Example
```
$ gogentype -json '{"str": "xxxx", "slice": [1,2,3], "sliceempty": [], "null": null, "int": 10, "zero": 0, "bool": true, "date": "2017-07-26T11:10:15+02:00", "obj": {"child":100}}' => generate newpkg_test.go
 >>
type TypeName1 struct {
    Str	        string      `json:"str"`
    Null        *string	    `json:"null"`
    Int	        int         `json:"int"`
    Zero        int         `json:"zero"`
    Bool        bool        `json:"bool"`
    Date        *time.Time  `json:"date"`
    Slice       []string    `json:"slice"`
    Sliceempty  []string    `json:"sliceempty"`
    Obj	        TypeName2   `json:"obj"`
}
type TypeName2 struct {
    Child       int         `json:"child"`
}
```
```
[sample.json]
{
  "url": "http://example.com/",
  "teachers": [
    { "id": 123, "name": "teacher aaa", "country": "Japan"},
    { "id": 124, "name": "teacher bbb", "country": "Netherlands"},
  ]
}

$ gogentype -file sample.json
 >>
type TypeName1 struct {
    Url       string     `json:"url"`
    Teachers  TypeName2  `json:"teachers"`
}
type TypeName2 struct {
    Id       int    `json:"id"`
    Name     string `json:"name"`
    Country  string `json:"country"`
}
```


# 6. gosubsrt [For Hobby]
Tweaking time lag of srt files.


## Installation
```
$ go get github.com/hiromaily/gotools/gosubsrt
```

## Usage
```
Options:
  -f  path of srt file.
  -t  time of tweaking duration.
e.g.:
  gosubsrt -f ./xxxxx.srt -t 1.5
```

## Example
```
prepare srt files somewhere
$ gosubsrt -f ./gosubsrt/srtfiles/sample.srt -t 6.2

[before]
1
00:00:10,950 --> 00:00:14,490
Ah! Fuck. Yeah, it doesn't matter.

[after]
1
00:00:17,150 --> 00:00:20,690
Ah! Fuck. Yeah, it doesn't matter.
```


# 7. godependency [WIP]
Create shell script file listed current commit id form outer packages like github.com directory.  
It has developed in progress yet.

## Installation
```
$ go get github.com/hiromaily/gotools/godependency
```

## Usage
```
Options:
  -target  path of github.com directory
e.g.:
  $ godepen -target ${HOME}/work/go/src/github.com
```

## Example
```
It shows latest checkout commit id from git directory for now

cd ${GOPATH}/src/github.com/treasure-data/td-client-go
git checkout 9cd8aa2ad7604fe42c622e83c753403754a5e729

cd ${GOPATH}/src/github.com/tylerb/graceful
git checkout d72b0151351a13d0421b763b88f791469c4f5dc7
```


# 8. gobulkdata [WIP]
Create CSV test dummy data

## Installation
```
$ go get github.com/hiromaily/gotools/gobulkdata
```

## Usage
```
Options:
  -f  File name.
  -t  File type.
e.g.:
  gobulkdata -f ${HOME}/work/go/src/github.com/hiromaily/gotools/text.txt -l 20
```


# 9. gochat [WIP]
chatting between client and server by TCP connection
