package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteVpcRouteRequest struct {
	// 路由ID

	RouteId string `json:"route_id"`
}

func (o DeleteVpcRouteRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteVpcRouteRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpcRouteRequest", string(data)}, " ")
}
