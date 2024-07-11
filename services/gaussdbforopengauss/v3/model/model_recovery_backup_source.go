package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecoveryBackupSource struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 恢复备份类型：backup，timestamp，different
	Type string `json:"type"`

	// 用于恢复的备份ID。
	BackupId *string `json:"backup_id,omitempty"`

	// UTC时间，时间戳
	RestoreTime *string `json:"restore_time,omitempty"`
}

func (o RecoveryBackupSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecoveryBackupSource struct{}"
	}

	return strings.Join([]string{"RecoveryBackupSource", string(data)}, " ")
}
