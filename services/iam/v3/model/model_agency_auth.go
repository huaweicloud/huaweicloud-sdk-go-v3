package model

import (
	"encoding/json"

	"strings"
)

//
type AgencyAuth struct {
	Identity *AgencyAuthIdentity `json:"identity"`
}

func (o AgencyAuth) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AgencyAuth struct{}"
	}

	return strings.Join([]string{"AgencyAuth", string(data)}, " ")
}
