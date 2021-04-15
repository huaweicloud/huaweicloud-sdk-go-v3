package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type NeutronCreateSecurityGroupResponse struct {
	SecurityGroup  *NeutronSecurityGroup `json:"security_group,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o NeutronCreateSecurityGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupResponse", string(data)}, " ")
}
