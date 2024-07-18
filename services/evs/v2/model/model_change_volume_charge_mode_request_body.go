package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeVolumeChargeModeRequestBody This is a auto create Body Object
type ChangeVolumeChargeModeRequestBody struct {

	// 要修改计费模式的云硬盘id列表，可以填写多个
	VolumeIds []string `json:"volume_ids"`

	// 云硬盘所挂载的虚拟机id, 如果volume_ids中没有多挂载的共享云硬盘，则此参数可无需填写。
	ServerId *string `json:"server_id,omitempty"`

	BssParam *BssParam2 `json:"bss_param,omitempty"`
}

func (o ChangeVolumeChargeModeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVolumeChargeModeRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeVolumeChargeModeRequestBody", string(data)}, " ")
}
