package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtendInfoReplication struct {
	// 目标副本ID

	DestinationBackupId *string `json:"destination_backup_id,omitempty"`
	// 目标还原点ID

	DestinationCheckpointId *string `json:"destination_checkpoint_id,omitempty"`
	// 目标project_id

	DestinationProjectId string `json:"destination_project_id"`
	// 目标区域

	DestinationRegion string `json:"destination_region"`
	// 源副本ID

	SourceBackupId string `json:"source_backup_id"`
	// 源还原点ID

	SourceCheckpointId *string `json:"source_checkpoint_id,omitempty"`
	// 源project_id

	SourceProjectId string `json:"source_project_id"`
	// 源区域

	SourceRegion string `json:"source_region"`
	// 源备份名称

	SourceBackupName *string `json:"source_backup_name,omitempty"`
	// 目标备份名称

	DestinationBackupName *string `json:"destination_backup_name,omitempty"`
}

func (o OpExtendInfoReplication) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoReplication struct{}"
	}

	return strings.Join([]string{"OpExtendInfoReplication", string(data)}, " ")
}
