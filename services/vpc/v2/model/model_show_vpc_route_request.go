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
type ShowVpcRouteRequest struct {
	RouteId string `json:"route_id"`
}

func (o ShowVpcRouteRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVpcRouteRequest", string(data)}, " ")
}
