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
type ShowVpcRouteResponse struct {
	Route *VpcRoute `json:"route,omitempty"`
}

func (o ShowVpcRouteResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVpcRouteResponse", string(data)}, " ")
}
