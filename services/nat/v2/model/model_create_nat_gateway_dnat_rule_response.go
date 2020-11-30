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
type CreateNatGatewayDnatRuleResponse struct {
	DnatRule       *NatGatewayDnatRuleResponseBody `json:"dnat_rule,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o CreateNatGatewayDnatRuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateNatGatewayDnatRuleResponse", string(data)}, " ")
}
