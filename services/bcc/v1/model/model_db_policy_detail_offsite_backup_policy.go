package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DbPolicyDetailOffsiteBackupPolicy struct {

	// 复制保留时长
	KeepDays *int32 `json:"keep_days,omitempty"`

	// 备份类型，增量或全量
	BackupType *string `json:"backup_type,omitempty"`

	// 目标区域ID
	DestinationRegion *string `json:"destination_region,omitempty"`

	// 目标项目ID
	DestinationProjectId *string `json:"destination_projectId,omitempty"`
}

func (o DbPolicyDetailOffsiteBackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbPolicyDetailOffsiteBackupPolicy struct{}"
	}

	return strings.Join([]string{"DbPolicyDetailOffsiteBackupPolicy", string(data)}, " ")
}
