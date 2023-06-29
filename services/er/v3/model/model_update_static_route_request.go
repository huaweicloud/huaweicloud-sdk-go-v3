package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStaticRouteRequest Request Object
type UpdateStaticRouteRequest struct {

	// 路由表ID
	RouteTableId string `json:"route_table_id"`

	// 路由ID
	RouteId string `json:"route_id"`

	Body *UpdateRouteRequestBody `json:"body,omitempty"`
}

func (o UpdateStaticRouteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStaticRouteRequest struct{}"
	}

	return strings.Join([]string{"UpdateStaticRouteRequest", string(data)}, " ")
}
