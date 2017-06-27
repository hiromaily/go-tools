package models

import (
	"fmt"
	"github.com/satori/go.uuid"
	u "github.com/hiromaily/golibs/utils"
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

func CreateRole() ([]byte, string, error) {
	var url = "http://localhost:3000/api/data/Role"

	d := Role{}
	d.Id = uuid.NewV4().String()

	incrementStr, err := getIncrement()
	if err != nil {
		fmt.Println(err)
	}
	d.Name = fmt.Sprintf("Role%s", incrementStr)

	roles := []string{"light", "full"}
	d.Type = u.PickOneFromEnum(roles)

	d.Description = nil
	d.OwnerId = nil
	d.Attr = nil
	d.Abilities = []string{}

	//to json
	data, err := convertJson(&d)

	return data, url, err
}
