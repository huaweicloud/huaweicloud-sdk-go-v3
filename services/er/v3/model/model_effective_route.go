package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EffectiveRoute 有效路由数据
type EffectiveRoute struct {

	// 路由ID
	RouteId *string `json:"route_id,omitempty"`

	// 路由目的地
	Destination *string `json:"destination,omitempty"`

	// 路由下一跳列表
	NextHops *[]RouteAttachment `json:"next_hops,omitempty"`

	// 是否黑洞路由
	IsBlackhole *bool `json:"is_blackhole,omitempty"`

	// 路由类型
	RouteType *string `json:"route_type,omitempty"`
}

func (o EffectiveRoute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EffectiveRoute struct{}"
	}

	return strings.Join([]string{"EffectiveRoute", string(data)}, " ")
}
