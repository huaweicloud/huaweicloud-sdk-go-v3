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
type CreateNatGatewayDnatRuleRequest struct {
	Body *CreateNatGatewayDnatRuleOption `json:"body,omitempty"`
}

func (o CreateNatGatewayDnatRuleRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateNatGatewayDnatRuleRequest", string(data)}, " ")
}
