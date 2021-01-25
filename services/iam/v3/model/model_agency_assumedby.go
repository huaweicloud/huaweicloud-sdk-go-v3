package model

import (
	"encoding/json"

	"strings"
)

//
type AgencyAssumedby struct {
	User *AgencyAssumedbyUser `json:"user"`
}

func (o AgencyAssumedby) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AgencyAssumedby struct{}"
	}

	return strings.Join([]string{"AgencyAssumedby", string(data)}, " ")
}
