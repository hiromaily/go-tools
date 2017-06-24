package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	ck "github.com/hiromaily/golibs/web/cookie"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"time"
	"os"
	"strconv"
)

type User struct {
	Id                  string     `json:"id"`
	Name                string     `json:"name"`
	FirstName           string     `json:"first_name"`
	LastName            string     `json:"last_name"`
	Password            string     `json:"password"`
	PasswordLastChanged *time.Time `json:"password_last_changed"`
	PasswordExpiration  *time.Time `json:"password_expiration"`
	PasswordExpired     bool       `json:"password_expired"`
	SendPasswordEmail   bool       `json:"send_password_email"`
	Email               string     `json:"email"`
	Gender              string     `json:"gender"`
	Type                string     `json:"type"`
	Enabled             bool       `json:"enabled"`
	DashboardId         *string    `json:"dashboard_id"`
	Comments            *string    `json:"comments"`
	Attr                *string    `json:"attr"`
	Departments         []string   `json:"departments"`
	Teams               []string   `json:"teams"`
	Roles               []string   `json:"roles"`
}

//curl 'http://localhost:3000/api/data/User?_dc=1498290488694'
// -H 'Cookie: gearbox=s%3AR8cD0biFGWAxk4lXCPahYMEXqdFT0k98.S8idxIcHAeqIBs7QdCwThfbIQ6DgSHj5p6b2D3xaHwU'
// -H 'Origin: http://localhost:3000'
// -H 'Accept-Encoding: gzip, deflate, br'
// -H 'Accept-Language: en-US,en;q=0.8,ja;q=0.6,nl;q=0.4,de;q=0.2,fr;q=0.2'
// -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36'
// -H 'Content-Type: application/json'
// -H 'Accept: */*'
// -H 'Referer: http://localhost:3000/'
// -H 'X-Requested-With: XMLHttpRequest'
// -H 'Connection: keep-alive'
// --data-binary '[{"id":"7974d21a-64aa-4b44-a550-3d830891cac9","name":"username","password_last_changed":null,"password_expiration":null,"attr":null,"type":"full","email":"email@test.com","gender":"Male","enabled":true,"comments":"Comment something","password":"H2&t3#I1","last_name":"lastname","first_name":"firstname","dashboard_id":null,"password_expired":false,"send_password_email":true,"departments":[],"teams":[],"roles":["55a6e956-2d61-4a68-91e0-d51b194dd700"]}]'
// --compressed

// --data-binary '[{
// "id":"7974d21a-64aa-4b44-a550-3d830891cac9",
// "name":"username",
// "first_name":"firstname",
// "last_name":"lastname",
// "password":"H2&t3#I1",
// "password_last_changed":null,
// "password_expiration":null,
// "password_expired":false,
// "send_password_email":true,
// "email":"email@test.com",
// "gender":"Male",
// "type":"full",
// "enabled":true,
// "dashboard_id":null,
// "comments":"Comment something",
// "attr":null,
// "departments":[],
// "teams":[],
// "roles":["55a6e956-2d61-4a68-91e0-d51b194dd700"]}]'

var (
	apiURL = "http://localhost:3000/api/data/User"
	domain = "localhost"
	cookieKey = "gearbox"

	firstNameBase = "first"
	lastNameBase = "last"
	increment = 1
)

func main() {
	//As precondition, cookie have to be stored on Chrome, so please access target url by Chrome first.

	//1. create test data
	user := createUser()

	//2. send post
	body, err := sendPost(user)
	if err != nil {
		fmt.Println(err)
	} else {
		//TODO: format
		fmt.Println(string(body))
	}
}

func getIncrement() (string, error){
	//increment = 1
	//os.O_APPEND
	filePath := os.Getenv("GOPATH") + "/src/github.com/hiromaily/gotools/goapigen/inclement"
	text, err := ioutil.ReadFile(filePath)
	if err != nil || len(text) == 0{
		//create or write file
		ioutil.WriteFile(filePath, []byte("1"), 0644)
		return "0001", fmt.Errorf("[ERROR] When calling `ioutil.ReadFile`: %v", err)
	}

	i, err := strconv.Atoi(string(text))
	if err != nil{
		//write file
		ioutil.WriteFile(filePath, []byte("1"), 0644)
		return "0001", fmt.Errorf("[ERROR] When calling `strconv.Atoi`: %v", err)
	}

	//write file
	i++
	ioutil.WriteFile(filePath, []byte(strconv.Itoa(i)), 0644)

	return fmt.Sprintf("%04d", i), nil
}

func createUser() *User {
	user := User{}
	user.Id = uuid.NewV4().String()

	incrementStr, err := getIncrement()
	if err != nil {
		fmt.Println(err)
	}

	user.Name = fmt.Sprintf("%s-%s%s", firstNameBase, lastNameBase, incrementStr)
	user.FirstName = fmt.Sprintf("%s%s", firstNameBase, incrementStr)
	user.LastName = fmt.Sprintf("%s%s", lastNameBase, incrementStr)

	user.Password = "H2&t3#I1"     //TODO:is it possible to send simple password to server directry??
	user.PasswordLastChanged = nil
	user.PasswordExpiration = nil
	user.PasswordExpired = false
	user.SendPasswordEmail = false
	user.Email = fmt.Sprintf("%s@api.test", user.Name)
	user.Gender = "Male"
	user.Type = "full"
	user.Enabled = true
	user.DashboardId = nil
	user.Comments = nil
	user.Attr = nil
	user.Departments = []string{}
	user.Teams = []string{}
	user.Roles = []string{"55a6e956-2d61-4a68-91e0-d51b194dd700"}

	return &user
}

func sendPost(user *User) ([]byte, error) {
	//create post data
	//values := url.Values{}
	//values.Set("token", token)
	//values.Add("device", device)

	//1. prepare data
	data, err := json.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] When calling `json.Marshal`: %v", err)
	}
	fmt.Println("[Debug] Json Data:", string(data))

	//b := new(bytes.Buffer)
	//json.NewEncoder(b).Encode(user)

	//2. prepare NewRequest data
	req, err := http.NewRequest(
		"POST",
		apiURL,
		//bytes.NewBuffer(jsonStr),
		bytes.NewReader(data),
	)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] When calling `http.NewRequest`: %v", err)
	}

	//3. get cookie
	cookie := ck.GetValue(domain, cookieKey)

	//4. set http header
	// Content-Type:application/json; charset=utf-8
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Cookie", fmt.Sprintf("%s=%s", cookieKey, cookie))

	//5. send
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] When calling `client.Do`: %v", err)
	}
	defer resp.Body.Close()

	//6. read response
	body, _ := ioutil.ReadAll(resp.Body)

	return body, err
}
