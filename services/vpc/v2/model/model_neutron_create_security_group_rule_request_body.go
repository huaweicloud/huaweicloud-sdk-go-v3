package model

import (
	"encoding/json"

	"strings"
)

//
type NeutronCreateSecurityGroupRuleRequestBody struct {
	SecurityGroupRule *NeutronCreateSecurityGroupRuleOption `json:"security_group_rule"`
}

func (o NeutronCreateSecurityGroupRuleRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupRuleRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupRuleRequestBody", string(data)}, " ")
}
