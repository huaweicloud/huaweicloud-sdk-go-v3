package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BaseDeleteRouteAction struct {

	// 路由子网
	Destination string `json:"destination"`

	// 下一跳id
	Nexthop string `json:"nexthop"`
}

func (o BaseDeleteRouteAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseDeleteRouteAction struct{}"
	}

	return strings.Join([]string{"BaseDeleteRouteAction", string(data)}, " ")
}
