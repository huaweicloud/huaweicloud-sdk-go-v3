package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BGP选路配置参数
type BgpOptions struct {

	// 配置路由在形成负载分担时不比较路由的AS-Path属性
	LoadBalancingAsPathIgnore *bool `json:"load_balancing_as_path_ignore,omitempty"`

	// 配置路由在形成负载分担时不比较相同长度的AS-Path属性
	LoadBalancingAsPathRelax *bool `json:"load_balancing_as_path_relax,omitempty"`
}

func (o BgpOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BgpOptions struct{}"
	}

	return strings.Join([]string{"BgpOptions", string(data)}, " ")
}
