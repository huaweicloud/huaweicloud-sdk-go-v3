/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type RejectVpcPeeringRequest struct {
	PeeringId string `json:"peering_id"`
}

func (o RejectVpcPeeringRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RejectVpcPeeringRequest", string(data)}, " ")
}
