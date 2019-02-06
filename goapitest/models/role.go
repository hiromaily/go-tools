package models

import (
	"fmt"
	u "github.com/hiromaily/golibs/utils"
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

func CreateRole() ([]byte, string, error) {
	var url = "http://localhost:3000/api/data/Role"

	d := Role{}
	ui, err := uuid.NewV4()
	if err != nil {
		return nil, "", err
	}
	d.Id = ui.String()

	incrementStr, err := getIncrement()
	if err != nil {
		return nil, "", err
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
