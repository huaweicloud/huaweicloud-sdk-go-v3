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
type DeleteNatGatewaySnatRuleResponse struct {
}

func (o DeleteNatGatewaySnatRuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteNatGatewaySnatRuleResponse", string(data)}, " ")
}
