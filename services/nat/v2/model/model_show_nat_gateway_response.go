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
type ShowNatGatewayResponse struct {
	NatGateway *NatGatewayResponseBody `json:"nat_gateway,omitempty"`
}

func (o ShowNatGatewayResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowNatGatewayResponse", string(data)}, " ")
}
