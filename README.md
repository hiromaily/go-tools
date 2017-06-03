# gotools

[![Go Report Card](https://goreportcard.com/badge/github.com/hiromaily/gotools)](https://goreportcard.com/report/github.com/hiromaily/gotools)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/hiromaily/gotools/master/LICENSE)

go tools

#1. gotestfile
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


#2. gocipher
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
  gcp -m e xxxxxxxx
    or
  gcp -m d xxxxxxxx
```


#3. godependency [WIP]
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


#4. gobulkdata [WIP]
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


#5. gochat [WIP]
chatting between client and server by TCP connection

