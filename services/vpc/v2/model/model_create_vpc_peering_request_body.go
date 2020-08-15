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
type CreateVpcPeeringRequestBody struct {
	Peering *CreateVpcPeeringOption `json:"peering"`
}

func (o CreateVpcPeeringRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateVpcPeeringRequestBody", string(data)}, " ")
}
