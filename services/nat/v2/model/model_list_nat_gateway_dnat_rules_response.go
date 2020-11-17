/*
 * NAT
 *
 * Open Api of Public Nat.
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListNatGatewayDnatRulesResponse struct {
	// 查询DNAT规则列表的响应体。
	DnatRules *[]NatGatewayDnatRuleResponseBody `json:"dnat_rules,omitempty"`
}

func (o ListNatGatewayDnatRulesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListNatGatewayDnatRulesResponse", string(data)}, " ")
}
