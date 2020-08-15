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
type ShowVpcPeeringResponse struct {
	Peering *VpcPeering `json:"peering,omitempty"`
}

func (o ShowVpcPeeringResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVpcPeeringResponse", string(data)}, " ")
}
