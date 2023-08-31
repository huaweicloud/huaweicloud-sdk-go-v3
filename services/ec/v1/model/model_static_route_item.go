package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StaticRouteItem struct {

	// 目标网络
	Prefix string `json:"prefix"`

	// 下一跳地址
	NextHop string `json:"next_hop"`

	// 接口名字
	InterfaceName string `json:"interface_name"`

	// 优先级
	Priority *int32 `json:"priority,omitempty"`

	// 自动检测
	TrackNqa bool `json:"track_nqa"`

	// 发布到企业连接网络
	PostToCloud bool `json:"post_to_cloud"`
}

func (o StaticRouteItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StaticRouteItem struct{}"
	}

	return strings.Join([]string{"StaticRouteItem", string(data)}, " ")
}
