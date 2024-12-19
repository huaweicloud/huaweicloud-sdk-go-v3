package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddGdgwRouteAction struct {

	// 路由子网
	Destination string `json:"destination"`

	// 下一跳id
	Nexthop string `json:"nexthop"`

	// 路由描述
	Description *string `json:"description,omitempty"`

	Type *RouteTypeOfGdgw `json:"type,omitempty"`
}

func (o AddGdgwRouteAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddGdgwRouteAction struct{}"
	}

	return strings.Join([]string{"AddGdgwRouteAction", string(data)}, " ")
}
