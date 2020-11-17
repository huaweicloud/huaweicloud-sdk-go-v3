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
type ShowNatGatewaySnatRuleRequest struct {
	SnatRuleId string `json:"snat_rule_id"`
}

func (o ShowNatGatewaySnatRuleRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowNatGatewaySnatRuleRequest", string(data)}, " ")
}
