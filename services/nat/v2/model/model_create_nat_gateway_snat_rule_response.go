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
type CreateNatGatewaySnatRuleResponse struct {
	SnatRule *NatGatewaySnatRuleResponseBody `json:"snat_rule,omitempty"`
}

func (o CreateNatGatewaySnatRuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateNatGatewaySnatRuleResponse", string(data)}, " ")
}
