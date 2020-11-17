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
type DeleteNatGatewayDnatRuleResponse struct {
}

func (o DeleteNatGatewayDnatRuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteNatGatewayDnatRuleResponse", string(data)}, " ")
}
