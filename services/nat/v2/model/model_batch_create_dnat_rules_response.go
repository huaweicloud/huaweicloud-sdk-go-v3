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
type BatchCreateDnatRulesResponse struct {
	// DNAT规则批量创建对象的响应体。
	DnatRules *[]NatGatewayDnatRuleResponseBody `json:"dnat_rules,omitempty"`
}

func (o BatchCreateDnatRulesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateDnatRulesResponse", string(data)}, " ")
}
