package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostRoute
type HostRoute struct {

	// 路由目的子网
	Destination *string `json:"destination,omitempty"`

	// 路由下一跳IP
	Nexthop *string `json:"nexthop,omitempty"`
}

func (o HostRoute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostRoute struct{}"
	}

	return strings.Join([]string{"HostRoute", string(data)}, " ")
}
