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
type DeleteVpcRouteRequest struct {
	RouteId string `json:"route_id"`
}

func (o DeleteVpcRouteRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteVpcRouteRequest", string(data)}, " ")
}
