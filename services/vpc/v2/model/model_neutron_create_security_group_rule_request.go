package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NeutronCreateSecurityGroupRuleRequest struct {
	Body *NeutronCreateSecurityGroupRuleRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateSecurityGroupRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupRuleRequest", string(data)}, " ")
}
