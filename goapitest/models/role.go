package models

import (
	"fmt"
	"github.com/satori/go.uuid"
)

//curl 'http://localhost:3000/api/data/Role?_dc=1498303478867'
// -H 'Cookie: gearbox=s%3AR8cD0biFGWAxk4lXCPahYMEXqdFT0k98.S8idxIcHAeqIBs7QdCwThfbIQ6DgSHj5p6b2D3xaHwU'
// -H 'Origin: http://localhost:3000'
// -H 'Accept-Encoding: gzip, deflate, br'
// -H 'Accept-Language: en-US,en;q=0.8,ja;q=0.6,nl;q=0.4,de;q=0.2,fr;q=0.2'
// -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36'
// -H 'Content-Type: application/json'
// -H 'Accept: */*' -H 'Referer: http://localhost:3000/'
// -H 'X-Requested-With: XMLHttpRequest'
// -H 'Connection: keep-alive'
// --data-binary '[{
//
// "id":"daef8658-3e58-43cb-8871-9752a00e241f",
// "abilities":[],
// "attr":null,
// "name":"newrole",
// "type":"light",
// "ownerid":null,
// "description":"aaaaa"}]'
// --compressed

type Role struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Description *string  `json:"description"`
	OwnerId     *string  `json:"ownerid"`
	Attr        *string  `json:"attr"`
	Abilities   []string `json:"abilities"`
}

func CreateRole(role string) ([]byte, string, error) {
	var url = "http://localhost:3000/api/data/Role"

	d := Role{}
	d.Id = uuid.NewV4().String()

	incrementStr, err := getIncrement()
	if err != nil {
		fmt.Println(err)
	}
	d.Name = fmt.Sprintf("Role%s", incrementStr)

	if role != "" {
		d.Type = role
	} else {
		d.Type = "full"
	}

	d.Description = nil
	d.OwnerId = nil
	d.Attr = nil
	d.Abilities = []string{}

	//to json
	data, err := convertJson(&d)

	return data, url, err
}
