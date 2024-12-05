package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateGdgwRoutetableRequestBodyGdgwRoutetable struct {
	AddRoutes *[]UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutes `json:"add_routes,omitempty"`

	// 删除的路由
	DelRoutes *[]UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutes `json:"del_routes,omitempty"`

	// 仅更新该条路由的附加信息，不执行交换机的路由更新操作。当前支持更新：路由描述-description信息
	UpdateRoutes *[]UpdateGdgwRoutetableRequestBodyGdgwRoutetableUpdateRoutes `json:"update_routes,omitempty"`
}

func (o UpdateGdgwRoutetableRequestBodyGdgwRoutetable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGdgwRoutetableRequestBodyGdgwRoutetable struct{}"
	}

	return strings.Join([]string{"UpdateGdgwRoutetableRequestBodyGdgwRoutetable", string(data)}, " ")
}
