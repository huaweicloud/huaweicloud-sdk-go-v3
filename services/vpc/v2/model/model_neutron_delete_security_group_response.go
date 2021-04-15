package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type NeutronDeleteSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteSecurityGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronDeleteSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteSecurityGroupResponse", string(data)}, " ")
}
