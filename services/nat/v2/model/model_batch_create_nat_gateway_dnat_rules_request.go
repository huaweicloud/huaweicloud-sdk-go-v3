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
type BatchCreateNatGatewayDnatRulesRequest struct {
	Body *BatchCreateNatGatewayDnatRulesRequestBody `json:"body,omitempty"`
}

func (o BatchCreateNatGatewayDnatRulesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateNatGatewayDnatRulesRequest", string(data)}, " ")
}
