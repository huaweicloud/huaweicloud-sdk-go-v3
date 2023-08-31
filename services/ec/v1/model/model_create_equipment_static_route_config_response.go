package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEquipmentStaticRouteConfigResponse Response Object
type CreateEquipmentStaticRouteConfigResponse struct {

	// 目标网络
	Prefix *string `json:"prefix,omitempty"`

	// 下一跳地址
	NextHop *string `json:"next_hop,omitempty"`

	// 接口名字
	InterfaceName *string `json:"interface_name,omitempty"`

	// 优先级
	Priority *int32 `json:"priority,omitempty"`

	// 自动检测
	TrackNqa *bool `json:"track_nqa,omitempty"`

	// 发布到企业连接网络
	PostToCloud    *bool `json:"post_to_cloud,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateEquipmentStaticRouteConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEquipmentStaticRouteConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateEquipmentStaticRouteConfigResponse", string(data)}, " ")
}
