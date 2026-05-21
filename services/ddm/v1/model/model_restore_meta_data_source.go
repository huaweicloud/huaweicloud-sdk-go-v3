package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreMetaDataSource struct {

	// 恢复时间。
	RestoreTime float32 `json:"restore_time,omitempty"`

	// 备份id。
	BackupId *string `json:"backup_id,omitempty"`
}

func (o RestoreMetaDataSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreMetaDataSource struct{}"
	}

	return strings.Join([]string{"RestoreMetaDataSource", string(data)}, " ")
}
