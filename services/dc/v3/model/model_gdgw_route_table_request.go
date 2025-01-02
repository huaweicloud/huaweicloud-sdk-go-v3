package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GdgwRouteTableRequest 修改路由入参对象
type GdgwRouteTableRequest struct {

	// 需要添加的路由
	AddRoutes *[]AddGdgwRouteAction `json:"add_routes,omitempty"`

	// 需要删除的路由
	DelRoutes *[]DeleteGdgwRouteAction `json:"del_routes,omitempty"`

	// 需要更新的路由 仅更新该条路由的附加信息，不执行交换机的路由更新操作。当前支持更新：路由描述-description信息
	UpdateRoutes *[]UpdateRouteAction `json:"update_routes,omitempty"`
}

func (o GdgwRouteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GdgwRouteTableRequest struct{}"
	}

	return strings.Join([]string{"GdgwRouteTableRequest", string(data)}, " ")
}
