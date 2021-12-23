package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type BackupReplicateRespBody struct {
	// 待复制的备份ID

	BackupId *string `json:"backup_id,omitempty"`
	// 复制的目标项目ID

	DestinationProjectId *string `json:"destination_project_id,omitempty"`
	// 复制的目标区域

	DestinationRegion *string `json:"destination_region,omitempty"`
	// 复制的目标区域存储库ID

	DestinationVaultId *string `json:"destination_vault_id,omitempty"`
	// 执行复制的项目ID

	ProjectId *string `json:"project_id,omitempty"`
	// 备份提供商ID，用于区分备份对象.

	ProviderId *string `json:"provider_id,omitempty"`
	// 复制记录ID

	ReplicationRecordId *string `json:"replication_record_id,omitempty"`
	// 复制的源区域

	SourceRegion *string `json:"source_region,omitempty"`
}

func (o BackupReplicateRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupReplicateRespBody struct{}"
	}

	return strings.Join([]string{"BackupReplicateRespBody", string(data)}, " ")
}
