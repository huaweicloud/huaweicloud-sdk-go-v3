package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RouteDestination 路由子网
type RouteDestination struct {
}

func (o RouteDestination) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteDestination struct{}"
	}

	return strings.Join([]string{"RouteDestination", string(data)}, " ")
}
