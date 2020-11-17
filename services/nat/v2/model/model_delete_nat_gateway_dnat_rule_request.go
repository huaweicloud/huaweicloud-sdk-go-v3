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

// Request Object
type DeleteNatGatewayDnatRuleRequest struct {
	NatGatewayId string `json:"nat_gateway_id"`
	DnatRuleId   string `json:"dnat_rule_id"`
}

func (o DeleteNatGatewayDnatRuleRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteNatGatewayDnatRuleRequest", string(data)}, " ")
}
