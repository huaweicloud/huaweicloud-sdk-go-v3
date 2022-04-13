package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckpointReplicateParam struct {
	// 本次复制是否自动触发，默认为false，代表手动触发

	AutoTrigger *bool `json:"auto_trigger,omitempty"`
	// 复制的目标项目ID

	DestinationProjectId string `json:"destination_project_id"`
	// 复制的目标区域id

	DestinationRegion string `json:"destination_region"`
	// 目标区域存储库ID

	DestinationVaultId string `json:"destination_vault_id"`
	// 跨区域复制时，是否启用加速从而缩短复制的时间，如果不指定，默认不启用加速，如果启用加速，会额外收取加速的费用。

	EnableAcceleration *bool `json:"enable_acceleration,omitempty"`
	// 存储库ID: uuid

	VaultId string `json:"vault_id"`
}

func (o CheckpointReplicateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointReplicateParam struct{}"
	}

	return strings.Join([]string{"CheckpointReplicateParam", string(data)}, " ")
}
