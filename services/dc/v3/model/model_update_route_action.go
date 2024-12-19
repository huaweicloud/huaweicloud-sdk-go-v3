package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRouteAction struct {

	// 路由子网
	Destination string `json:"destination"`

	// 下一跳id
	Nexthop string `json:"nexthop"`

	// 路由描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateRouteAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRouteAction struct{}"
	}

	return strings.Join([]string{"UpdateRouteAction", string(data)}, " ")
}
