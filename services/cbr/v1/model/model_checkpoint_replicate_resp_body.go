package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckpointReplicateRespBody struct {

	// 待复制的备份列表
	Backups []CheckpointReplicateRespbackups `json:"backups" xml:"backups"`

	// 复制的目标项目ID
	DestinationProjectId string `json:"destination_project_id" xml:"destination_project_id"`

	// 复制的目标区域
	DestinationRegion string `json:"destination_region" xml:"destination_region"`

	// 目标区域存储库ID
	DestinationVaultId string `json:"destination_vault_id" xml:"destination_vault_id"`

	// 执行复制的项目ID
	ProjectId string `json:"project_id" xml:"project_id"`

	// 备份提供商ID
	ProviderId string `json:"provider_id" xml:"provider_id"`

	// 复制的源区域
	SourceRegion string `json:"source_region" xml:"source_region"`

	// 存储库ID
	VaultId string `json:"vault_id" xml:"vault_id"`
}

func (o CheckpointReplicateRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointReplicateRespBody struct{}"
	}

	return strings.Join([]string{"CheckpointReplicateRespBody", string(data)}, " ")
}
