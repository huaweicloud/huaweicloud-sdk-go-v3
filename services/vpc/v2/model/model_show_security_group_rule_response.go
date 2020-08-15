/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type ShowSecurityGroupRuleResponse struct {
	SecurityGroupRule *SecurityGroupRule `json:"security_group_rule,omitempty"`
}

func (o ShowSecurityGroupRuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowSecurityGroupRuleResponse", string(data)}, " ")
}
