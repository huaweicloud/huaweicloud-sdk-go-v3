package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ips特性状态返回查询
type IpsSwitchResponseDto struct {

	// object_id
	ObjectId *string `json:"object_id,omitempty"`

	// 基础防御状态
	BasicDefenseStatus *int32 `json:"basic_defense_status,omitempty"`

	// 虚拟补丁状态
	VirtualPatchesStauts *int32 `json:"virtual_patches_stauts,omitempty"`
}

func (o IpsSwitchResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsSwitchResponseDto struct{}"
	}

	return strings.Join([]string{"IpsSwitchResponseDto", string(data)}, " ")
}
