package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type NeutronShowSecurityGroupResponse struct {
	SecurityGroup  *NeutronSecurityGroup `json:"security_group,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o NeutronShowSecurityGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronShowSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowSecurityGroupResponse", string(data)}, " ")
}
