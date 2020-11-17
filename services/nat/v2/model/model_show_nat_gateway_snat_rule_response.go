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
type ShowNatGatewaySnatRuleResponse struct {
	SnatRule *NatGatewaySnatRuleResponseBody `json:"snat_rule,omitempty"`
}

func (o ShowNatGatewaySnatRuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowNatGatewaySnatRuleResponse", string(data)}, " ")
}
