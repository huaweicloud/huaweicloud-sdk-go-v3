package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NeutronCreateSecurityGroupRequest struct {
	Body *NeutronCreateSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateSecurityGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupRequest", string(data)}, " ")
}
