package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpsSwitchResponseDto ips特性状态返回查询
type IpsSwitchResponseDto struct {

	// ips开关id
	Id *string `json:"id,omitempty"`

	// 基础防御状态
	BasicDefenseStatus *int32 `json:"basic_defense_status,omitempty"`

	// 虚拟补丁状态
	VirtualPatchesStatus *int32 `json:"virtual_patches_status,omitempty"`
}

func (o IpsSwitchResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsSwitchResponseDto struct{}"
	}

	return strings.Join([]string{"IpsSwitchResponseDto", string(data)}, " ")
}
