package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteGdgwRouteAction struct {

	// 路由子网
	Destination string `json:"destination"`

	// 下一跳id
	Nexthop string `json:"nexthop"`

	Type *RouteTypeOfGdgw `json:"type,omitempty"`
}

func (o DeleteGdgwRouteAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGdgwRouteAction struct{}"
	}

	return strings.Join([]string{"DeleteGdgwRouteAction", string(data)}, " ")
}
