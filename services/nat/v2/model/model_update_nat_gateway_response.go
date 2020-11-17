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
type UpdateNatGatewayResponse struct {
	NatGateway *NatGatewayResponseBody `json:"nat_gateway,omitempty"`
}

func (o UpdateNatGatewayResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateNatGatewayResponse", string(data)}, " ")
}
