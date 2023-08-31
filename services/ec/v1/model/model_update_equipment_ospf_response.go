package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEquipmentOspfResponse Response Object
type UpdateEquipmentOspfResponse struct {

	// 是否启用OSPF
	OspfEnabled *bool `json:"ospf_enabled,omitempty"`

	// 区域标识
	AreaId *int64 `json:"area_id,omitempty"`

	// 发布到企业连接网络
	PostToCloud *bool `json:"post_to_cloud,omitempty"`

	// 发送Hello报文的时间间隔，单位是秒
	HelloTimer *int32 `json:"hello_timer,omitempty"`

	// 点分十进制格式，OSPF协议使用全网唯一的Router ID
	RouterId *string `json:"router_id,omitempty"`

	// 启用OSPF协议的接口列表
	Interfaces *[]string `json:"interfaces,omitempty"`

	// 是否启用前缀过滤
	FilterEnabled *bool `json:"filter_enabled,omitempty"`

	// 白名单列表
	TrustList *[]string `json:"trust_list,omitempty"`

	// 黑名单列表
	BlockList      *[]string `json:"block_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateEquipmentOspfResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEquipmentOspfResponse struct{}"
	}

	return strings.Join([]string{"UpdateEquipmentOspfResponse", string(data)}, " ")
}
