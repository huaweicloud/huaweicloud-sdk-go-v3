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
type BatchCreateNatGatewayDnatRulesResponse struct {
	// DNAT规则批量创建对象的响应体。
	DnatRules      *[]NatGatewayDnatRuleResponseBody `json:"dnat_rules,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o BatchCreateNatGatewayDnatRulesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateNatGatewayDnatRulesResponse", string(data)}, " ")
}
