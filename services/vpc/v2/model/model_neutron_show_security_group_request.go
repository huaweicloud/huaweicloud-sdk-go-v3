package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NeutronShowSecurityGroupRequest struct {
	SecurityGroupId string `json:"security_group_id"`
}

func (o NeutronShowSecurityGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronShowSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowSecurityGroupRequest", string(data)}, " ")
}
