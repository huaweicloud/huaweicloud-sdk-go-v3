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
type UpdateVpcPeeringRequest struct {
	PeeringId string                       `json:"peering_id"`
	Body      *UpdateVpcPeeringRequestBody `json:"body,omitempty"`
}

func (o UpdateVpcPeeringRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateVpcPeeringRequest", string(data)}, " ")
}
