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
type UpdateNatGatewayDnatRuleRequest struct {
	DnatRuleId string                               `json:"dnat_rule_id"`
	Body       *UpdateNatGatewayDnatRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateNatGatewayDnatRuleRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateNatGatewayDnatRuleRequest", string(data)}, " ")
}
