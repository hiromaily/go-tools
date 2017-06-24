package models

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

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

var (
	firstNameBase = "first"
	lastNameBase  = "last"
)

func CreateUser() ([]byte, string, error) {
	var url = "http://localhost:3000/api/data/User"

	d := User{}
	d.Id = uuid.NewV4().String()

	incrementStr, err := getIncrement()
	if err != nil {
		fmt.Println(err)
	}

	d.Name = fmt.Sprintf("%s-%s%s", firstNameBase, lastNameBase, incrementStr)
	d.FirstName = fmt.Sprintf("%s%s", firstNameBase, incrementStr)
	d.LastName = fmt.Sprintf("%s%s", lastNameBase, incrementStr)

	d.Password = "H2&t3#I1" //TODO:is it possible to send simple password to server directry??
	d.PasswordLastChanged = nil
	d.PasswordExpiration = nil
	d.PasswordExpired = false
	d.SendPasswordEmail = false
	d.Email = fmt.Sprintf("%s@api.test", d.Name)
	d.Gender = "Male"
	d.Type = "full"
	d.Enabled = true
	d.DashboardId = nil
	d.Comments = nil
	d.Attr = nil
	d.Departments = []string{}
	d.Teams = []string{}
	d.Roles = []string{"55a6e956-2d61-4a68-91e0-d51b194dd700"}

	//to json
	data, err := convertJson(&d)

	return data, url, err
}
