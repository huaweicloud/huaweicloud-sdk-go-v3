package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新路由表路由对象动作，可选add、mod、del
type RouteTableRouteAction struct {

	// 新增路由条目，type，destination，nexthop必选
	Add *[]AddRouteTableRoute `json:"add,omitempty"`

	// 修改路由条目，type，destination，nexthop必选
	Mod *[]ModRouteTableRoute `json:"mod,omitempty"`

	// 删除路由条目，destination必选
	Del *[]DelRouteTableRoute `json:"del,omitempty"`
}

func (o RouteTableRouteAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteTableRouteAction struct{}"
	}

	return strings.Join([]string{"RouteTableRouteAction", string(data)}, " ")
}
