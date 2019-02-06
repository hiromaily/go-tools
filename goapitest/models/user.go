package models

import (
	"fmt"
	u "github.com/hiromaily/golibs/utils"
	"github.com/icrowley/fake"
	"github.com/satori/go.uuid"
	"time"
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

var (
	firstNameBase = "first"
	lastNameBase  = "last"
)

func CreateUser() ([]byte, string, error) {
	var url = "http://localhost:3000/api/data/User"

	d := User{}
	ui, err := uuid.NewV4()
	if err != nil {
		return nil, "", err
	}
	d.Id = ui.String()

	incrementStr, err := getIncrement()
	if err != nil {
		fmt.Println(err)
	}

	d.Name = fmt.Sprintf("%s-%s%s", firstNameBase, lastNameBase, incrementStr)
	//d.FirstName = fmt.Sprintf("%s%s", firstNameBase, incrementStr)
	//d.LastName = fmt.Sprintf("%s%s", lastNameBase, incrementStr)
	d.FirstName = fake.FirstName()
	d.LastName = fake.LastName()

	d.Password = "H2&t3#I1" //TODO:is it possible to send simple password to server directry?? => No!
	//d.Password = fake.SimplePassword()

	d.PasswordLastChanged = nil
	d.PasswordExpiration = nil
	d.PasswordExpired = false
	d.SendPasswordEmail = false
	d.Email = fmt.Sprintf("%s@api.test", d.Name)
	d.Gender = "Male"

	roles := []string{"light", "full"}
	d.Type = u.PickOneFromEnum(roles)

	d.Enabled = true
	d.DashboardId = nil
	d.Comments = nil
	d.Attr = nil
	d.Departments = []string{}
	d.Teams = []string{}
	if d.Type == "full" {
		d.Roles = []string{"55a6e956-2d61-4a68-91e0-d51b194dd700"}
	}

	//to json
	data, err := convertJson(&d)

	return data, url, err
}
