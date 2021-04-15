package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NeutronDeleteSecurityGroupRuleRequest struct {
	SecurityGroupRuleId string `json:"security_group_rule_id"`
}

func (o NeutronDeleteSecurityGroupRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronDeleteSecurityGroupRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteSecurityGroupRuleRequest", string(data)}, " ")
}
