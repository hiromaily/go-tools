# gotools

[![Go Report Card](https://goreportcard.com/badge/github.com/hiromaily/gotools)](https://goreportcard.com/report/github.com/hiromaily/gotools)
[![codebeat badge](https://codebeat.co/badges/4e4e3273-f177-4b4a-812d-b902c01d8b3d)](https://codebeat.co/projects/github-com-hiromaily-go-tools-master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/6b57918b3861422b8d67add70fe0bb59)](https://www.codacy.com/app/hiromaily2/go-tools?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=hiromaily/go-tools&amp;utm_campaign=Badge_Grade)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/hiromaily/gotools/master/LICENSE)

go tools


# 1. encryption
Create encrypted string both (encode and decode)
It requires below environment variable, e.g. 
```
export 'ENC_KEY=M#XF#R+XuKF4J_<'
export 'ENC_IV=@~nL-3OlQ<c2y@pA'
```

## Usage
```
Options:
  -m  e: encode, d: decode.
e.g.:
  encryption -m e xxxxxxxx
    or
  encryption -m d xxxxxxxx
```

## Example
```
$ encryption -m e secret_string
 => gtBl3kNqSAJGvJjnvUU9HQ==

$ encryption -m d gtBl3kNqSAJGvJjnvUU9HQ==
 => secret_string
```


# 2. cookie [For Development]
Retrieve cookie data by domain from chrome.  

## Usage
```
e.g.:
  cookie github.com
```

## Example
```
$ cookie localhost
 => localhost/key: value
```


# 3. gen-tls-cert [For Development]
Generate TLS certificate files using golang package.  

## Usage
```
e.g.:
  gen-tls-cert -host hy
```

## Example
```
$ gen-tls-cert -host hy
 => 
2017/06/28 21:06:43 written cert.pem
2017/06/28 21:06:43 written key.pem
```


# 4. go-testfile [For Golang Development]
Create template file of xxx_test.go

## Usage
```
Usage: go-testfile [options...]

Options:
  -n     package name

e.g.
 $ go-testfile -n new-package-name
```

## Example
```
$ go-testfile -n newpkg
 => generate newpkg_test.go file
```


# 5. gen-struct [For Golang Development]
Create golang type struct from json data.

## Usage
```
Usage: gogentype [options...]
Options:
  -json  Package name.
e.g.:
  gen-struct -json '{"str": "xxxx", "slice": [1,2,3], "sliceempty": [], "null": null, "int": 10, "zero": 0, "bool": true, "obj": {"child":100}}'
 or
  gen-struct -file sample.json

Note:null value can not be detected proper type.
```

## Example
```
$ gen-struct -json '{"str": "xxxx", "slice": [1,2,3], "sliceempty": [], "null": null, "int": 10, "zero": 0, "bool": true, "date": "2017-07-26T11:10:15+02:00", "obj": {"child":100}}' => generate newpkg_test.go
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


# 6. subsrt [For Hobby]
Tweaking time lag of srt files.

## Usage
```
Options:
  -f  path of srt file.
  -t  time of tweaking duration.
e.g.:
  subsrt -f ./xxxxx.srt -t 1.5
```

## Example
```
prepare srt files somewhere
$ subsrt -f ./gosubsrt/srtfiles/sample.srt -t 6.2

[before]
1
00:00:10,950 --> 00:00:14,490
Ah! Fuck. Yeah, it doesn't matter.

[after]
1
00:00:17,150 --> 00:00:20,690
Ah! Fuck. Yeah, it doesn't matter.
```


# 7. go-dependency [WIP]
Create shell script file listed current commit id form outer packages like github.com directory.  
It has developed in progress yet.

## Usage
```
Options:
  -target  path of github.com directory
e.g.:
  $ go-dependency -target ${HOME}/work/go/src/github.com
```

## Example
```
It shows latest checkout commit id from git directory for now

cd ${GOPATH}/src/github.com/treasure-data/td-client-go
git checkout 9cd8aa2ad7604fe42c622e83c753403754a5e729

cd ${GOPATH}/src/github.com/tylerb/graceful
git checkout d72b0151351a13d0421b763b88f791469c4f5dc7
```


# 8. gochat [WIP]
chatting between client and server by TCP connection
