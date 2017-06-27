package models

import (
	"fmt"
	"github.com/satori/go.uuid"
)

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
