package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStaticRouteRequest Request Object
type ShowStaticRouteRequest struct {

	// 路由表ID
	RouteTableId string `json:"route_table_id"`

	// 路由ID
	RouteId string `json:"route_id"`
}

func (o ShowStaticRouteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStaticRouteRequest struct{}"
	}

	return strings.Join([]string{"ShowStaticRouteRequest", string(data)}, " ")
}
