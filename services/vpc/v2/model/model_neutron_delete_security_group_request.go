package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NeutronDeleteSecurityGroupRequest struct {
	SecurityGroupId string `json:"security_group_id"`
}

func (o NeutronDeleteSecurityGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronDeleteSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteSecurityGroupRequest", string(data)}, " ")
}
