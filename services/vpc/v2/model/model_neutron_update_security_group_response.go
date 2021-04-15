package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type NeutronUpdateSecurityGroupResponse struct {
	SecurityGroup  *NeutronSecurityGroup `json:"security_group,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o NeutronUpdateSecurityGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronUpdateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronUpdateSecurityGroupResponse", string(data)}, " ")
}
