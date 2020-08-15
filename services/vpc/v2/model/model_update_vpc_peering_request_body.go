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

//
type UpdateVpcPeeringRequestBody struct {
	Peering *UpdateVpcPeeringOption `json:"peering"`
}

func (o UpdateVpcPeeringRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateVpcPeeringRequestBody", string(data)}, " ")
}
