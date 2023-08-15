package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStaticRouteRequest Request Object
type DeleteStaticRouteRequest struct {

	// 路由表ID
	RouteTableId string `json:"route_table_id"`

	// 路由ID
	RouteId string `json:"route_id"`
}

func (o DeleteStaticRouteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStaticRouteRequest struct{}"
	}

	return strings.Join([]string{"DeleteStaticRouteRequest", string(data)}, " ")
}
