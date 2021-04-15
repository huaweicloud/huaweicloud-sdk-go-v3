package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type NeutronDeleteSecurityGroupRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteSecurityGroupRuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronDeleteSecurityGroupRuleResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteSecurityGroupRuleResponse", string(data)}, " ")
}
