package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NeutronUpdateSecurityGroupRequest struct {
	SecurityGroupId string `json:"security_group_id"`

	Body *NeutronUpdateSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o NeutronUpdateSecurityGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronUpdateSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronUpdateSecurityGroupRequest", string(data)}, " ")
}
