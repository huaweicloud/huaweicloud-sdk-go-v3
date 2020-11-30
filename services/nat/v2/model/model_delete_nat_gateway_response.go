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
type DeleteNatGatewayResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNatGatewayResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteNatGatewayResponse", string(data)}, " ")
}
