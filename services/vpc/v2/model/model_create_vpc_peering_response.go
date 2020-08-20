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

// Response Object
type CreateVpcPeeringResponse struct {
	Peering *VpcPeering `json:"peering,omitempty"`
}

func (o CreateVpcPeeringResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateVpcPeeringResponse", string(data)}, " ")
}
