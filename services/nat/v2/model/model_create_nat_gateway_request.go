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
type CreateNatGatewayRequest struct {
	Body *CreateNatGatewayRequestBody `json:"body,omitempty"`
}

func (o CreateNatGatewayRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateNatGatewayRequest", string(data)}, " ")
}
