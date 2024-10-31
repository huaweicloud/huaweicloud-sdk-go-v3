package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpsSwitchResponseDto ips特性状态返回查询
type IpsSwitchResponseDto struct {

	// ips开关id，此处为互联网边界防护对象id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得
	Id *string `json:"id,omitempty"`

	// 基础防御状态，0表示关闭，1表示开启
	BasicDefenseStatus *int32 `json:"basic_defense_status,omitempty"`

	// 虚拟补丁状态，0表示关闭，1表示开启
	VirtualPatchesStatus *int32 `json:"virtual_patches_status,omitempty"`
}

func (o IpsSwitchResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsSwitchResponseDto struct{}"
	}

	return strings.Join([]string{"IpsSwitchResponseDto", string(data)}, " ")
}
