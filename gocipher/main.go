package main

import (
	"flag"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	enc "github.com/hiromaily/golibs/cipher/encryption"
	"os"
)

var (
	mode = flag.String("m", "e", "e:encode, d:decode")
)

var usage = `Usage: %s [options...]
Options:
  -m  e:encode, d:decode.
e.g.:
  gcp -m e xxxxxxxx
    or
  gcp -m d xxxxxxxx
`

type Params struct {
	Name      string
	Uppercase string
}

func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[GOTOOLS GOTEST]", "/var/log/go/gotool.log")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, os.Args[0]))
	}

	flag.Parse()

	//first argument
	//for i,val := range os.Args{
	//	fmt.Printf("%d:%s\n", i, val)
	//}

	if len(os.Args) != 4{
		flag.Usage()
		os.Exit(1)
		return
	}
}

func setup(){
	size := 16
	key := os.Getenv("ENC_KEY")
	iv := os.Getenv("ENC_IV")

	if key == "" || iv == ""{
		fmt.Errorf("%s", "set Environment Valuable: ENC_KEY, ENC_IV")
		os.Exit(1)
	}

	enc.NewCrypt(size, key, iv)
}
func main() {
	setup()

	crypt := enc.GetCryptInstance()
	targetStr := os.Args[3]
	fmt.Printf("target string is %s\n", targetStr)

	switch *mode {
	case "e":
		//encode
		fmt.Println(crypt.EncryptBase64(targetStr))
	case "d":
		//decode
		fmt.Println(crypt.DecryptBase64(targetStr))
	default:
		fmt.Errorf("%s", "arguments is wrong.")
	}
}
