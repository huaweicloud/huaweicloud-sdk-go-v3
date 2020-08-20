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
type CreateVpcRouteResponse struct {
	Route *VpcRoute `json:"route,omitempty"`
}

func (o CreateVpcRouteResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateVpcRouteResponse", string(data)}, " ")
}
