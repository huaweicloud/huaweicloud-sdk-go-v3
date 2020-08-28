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

// Request Object
type ShowSecurityGroupRuleRequest struct {
	SecurityGroupRuleId string `json:"security_group_rule_id"`
}

func (o ShowSecurityGroupRuleRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowSecurityGroupRuleRequest", string(data)}, " ")
}
