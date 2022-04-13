package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupRestore struct {
	// 恢复的映射关系(整机恢复时必填，卷恢复时可选但是不会用到填写的值）

	Mappings *[]BackupRestoreServerMapping `json:"mappings,omitempty"`
	// 恢复后是否开始，默认开机。

	PowerOn *bool `json:"power_on,omitempty"`
	// 恢复的目标虚拟机ID（整机恢复时必填）

	ServerId *string `json:"server_id,omitempty"`
	// 恢复的目标卷ID（卷恢复时必填）

	VolumeId *string `json:"volume_id,omitempty"`
	// 待恢复的目标资源ID

	ResourceId *string `json:"resource_id,omitempty"`
}

func (o BackupRestore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupRestore struct{}"
	}

	return strings.Join([]string{"BackupRestore", string(data)}, " ")
}
