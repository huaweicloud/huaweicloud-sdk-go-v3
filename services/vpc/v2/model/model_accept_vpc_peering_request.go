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
type AcceptVpcPeeringRequest struct {
	PeeringId string `json:"peering_id"`
}

func (o AcceptVpcPeeringRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AcceptVpcPeeringRequest", string(data)}, " ")
}
