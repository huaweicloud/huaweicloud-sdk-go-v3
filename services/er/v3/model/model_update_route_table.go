package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新路由表请求体
type UpdateRouteTable struct {

	// 路由器表名称
	Name *string `json:"name,omitempty"`

	// 路由器表描述信息
	Description *string `json:"description,omitempty"`

	// 是否选路忽略as path
	LoadBalancingAsPathIgnore *bool `json:"load_balancing_as_path_ignore,omitempty"`

	BgpOptions *BgpOptions `json:"bgp_options,omitempty"`
}

func (o UpdateRouteTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRouteTable struct{}"
	}

	return strings.Join([]string{"UpdateRouteTable", string(data)}, " ")
}
