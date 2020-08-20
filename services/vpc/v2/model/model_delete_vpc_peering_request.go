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
type DeleteVpcPeeringRequest struct {
	PeeringId string `json:"peering_id"`
}

func (o DeleteVpcPeeringRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteVpcPeeringRequest", string(data)}, " ")
}
