package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type BackupReplicateRespBody struct {

	// 待复制的备份ID
	BackupId *string `json:"backup_id,omitempty" xml:"backup_id"`

	// 复制的目标项目ID
	DestinationProjectId *string `json:"destination_project_id,omitempty" xml:"destination_project_id"`

	// 复制的目标区域
	DestinationRegion *string `json:"destination_region,omitempty" xml:"destination_region"`

	// 复制的目标区域存储库ID
	DestinationVaultId *string `json:"destination_vault_id,omitempty" xml:"destination_vault_id"`

	// 执行复制的项目ID
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 资源类型id
	ProviderId *string `json:"provider_id,omitempty" xml:"provider_id"`

	// 复制记录ID
	ReplicationRecordId *string `json:"replication_record_id,omitempty" xml:"replication_record_id"`

	// 复制的源区域
	SourceRegion *string `json:"source_region,omitempty" xml:"source_region"`
}

func (o BackupReplicateRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupReplicateRespBody struct{}"
	}

	return strings.Join([]string{"BackupReplicateRespBody", string(data)}, " ")
}
