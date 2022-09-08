package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 路由表
type CreateRouteTable struct {

	// 路由器表名称
	Name string `json:"name"`

	// 路由器表描述信息
	Description *string `json:"description,omitempty"`

	// 是否选路忽略as path
	LoadBalancingAsPathIgnore *bool `json:"load_balancing_as_path_ignore,omitempty"`

	BgpOptions *BgpOptions `json:"bgp_options,omitempty"`

	// 标签信息
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateRouteTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRouteTable struct{}"
	}

	return strings.Join([]string{"CreateRouteTable", string(data)}, " ")
}
