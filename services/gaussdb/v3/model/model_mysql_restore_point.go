package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlRestorePoint struct {

	// PITR。要恢复的时间点。
	RestoreTime *int64 `json:"restore_time,omitempty"`

	// 源实例ID。
	SourceInstanceId string `json:"source_instance_id"`

	// 备份文件ID。
	BackupId *string `json:"backup_id,omitempty"`

	// 备份类型。当参数为空时，backup_id不能为空，即默认按备份文件恢复。 当参数不为空时，取值范围： - backup：表示按备份文件恢复； - timestamp：表示按时间点恢复；
	Type *string `json:"type,omitempty"`
}

func (o MysqlRestorePoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlRestorePoint struct{}"
	}

	return strings.Join([]string{"MysqlRestorePoint", string(data)}, " ")
}
