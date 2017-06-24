package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var increment = 1

func getIncrement() (string, error) {
	//increment = 1
	//os.O_APPEND
	filePath := os.Getenv("GOPATH") + "/src/github.com/hiromaily/gotools/goapitest/inclement"
	text, err := ioutil.ReadFile(filePath)
	if err != nil || len(text) == 0 {
		//create or write file
		ioutil.WriteFile(filePath, []byte("1"), 0644)
		return "0001", fmt.Errorf("[ERROR] When calling `ioutil.ReadFile`: %v", err)
	}

	i, err := strconv.Atoi(string(text))
	if err != nil {
		//write file
		ioutil.WriteFile(filePath, []byte("1"), 0644)
		return "0001", fmt.Errorf("[ERROR] When calling `strconv.Atoi`: %v", err)
	}

	//write file
	i++
	ioutil.WriteFile(filePath, []byte(strconv.Itoa(i)), 0644)

	return fmt.Sprintf("%04d", i), nil
}

func convertJson(model interface{}) ([]byte, error) {
	data, err := json.Marshal(model)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] When calling `json.Marshal`: %v\n", err)
	}
	fmt.Println("[Debug] Json Data:", string(data))

	//b := new(bytes.Buffer)
	//json.NewEncoder(b).Encode(user)

	return data, nil
}
