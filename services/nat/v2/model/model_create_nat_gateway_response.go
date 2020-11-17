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
type CreateNatGatewayResponse struct {
	NatGateway *NatGatewayResponseBody `json:"nat_gateway,omitempty"`
}

func (o CreateNatGatewayResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateNatGatewayResponse", string(data)}, " ")
}
