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
type ShowVpcPeeringRequest struct {
	PeeringId string `json:"peering_id"`
}

func (o ShowVpcPeeringRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVpcPeeringRequest", string(data)}, " ")
}
