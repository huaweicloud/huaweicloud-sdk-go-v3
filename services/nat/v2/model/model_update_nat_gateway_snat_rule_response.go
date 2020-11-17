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
type UpdateNatGatewaySnatRuleResponse struct {
	SnatRule *NatGatewaySnatRuleResponseBody `json:"snat_rule,omitempty"`
}

func (o UpdateNatGatewaySnatRuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateNatGatewaySnatRuleResponse", string(data)}, " ")
}
