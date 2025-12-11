package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecycleBackupV3 struct {

	// 备份级别。
	BackupLevel *string `json:"backup_level,omitempty"`

	// 备份ID。
	BackupId *string `json:"backup_id,omitempty"`

	// 备份名称。
	BackupName *string `json:"backup_name,omitempty"`

	// 备份大小，单位是字节。
	Size *int32 `json:"size,omitempty"`

	// 回收状态。
	Status *string `json:"status,omitempty"`

	// 备份开始时间。
	BeginTime *string `json:"begin_time,omitempty"`

	// 备份结束时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o RecycleBackupV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleBackupV3 struct{}"
	}

	return strings.Join([]string{"RecycleBackupV3", string(data)}, " ")
}
