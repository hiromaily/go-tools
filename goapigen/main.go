package main

import (
	"bytes"
	"flag"
	"fmt"
	ck "github.com/hiromaily/golibs/web/cookie"
	"github.com/hiromaily/gotools/goapigen/models"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	API       = flag.String("m", "", "API Name")
	domain    = "localhost"
	cookieKey = "gearbox"
)

var usage = `Usage: %s [options...]
Options:
  -m  API Name.
Models:
  user: user api
  role: role api
e.g.:
  goapigen -m user

`

func init() {
	flag.Parse()

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, os.Args[0]))
	}

	if *API == "" {
		flag.Usage()

		os.Exit(1)
		return
	}
}

func main() {
	//As precondition, cookie have to be stored on Chrome, so please access target url by Chrome first.
	var (
		data []byte
		url  string
		err  error
	)

	//1. mode
	switch *API {
	case "user":
		data, url, err = models.CreateUser()
	case "role":
		data, url, err = models.CreateRole("light")
	default:
		err = fmt.Errorf("[ERROR] API arguments is invalid.")
		flag.Usage()
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	//2. send post
	body, err := sendPost(data, url)
	if err != nil {
		fmt.Println(err)
	} else {
		//TODO: format
		fmt.Println(string(body))
	}
}

func sendPost(data []byte, url string) ([]byte, error) {

	//1. prepare NewRequest data
	req, err := http.NewRequest(
		"POST",
		url,
		//bytes.NewBuffer(jsonStr),
		bytes.NewReader(data),
	)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] When calling `http.NewRequest`: %v", err)
	}

	//2. get cookie
	cookie := ck.GetValue(domain, cookieKey)

	//3. set http header
	// Content-Type:application/json; charset=utf-8
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Cookie", fmt.Sprintf("%s=%s", cookieKey, cookie))

	//4. send
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] When calling `client.Do`: %v", err)
	}
	defer resp.Body.Close()

	//5. read response
	body, _ := ioutil.ReadAll(resp.Body)

	return body, err
}
