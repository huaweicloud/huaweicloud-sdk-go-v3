package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupJobEndpointInfo 备份迁移任务恢复目标RDS for SQL Server实例信息。
type BackupJobEndpointInfo struct {

	// 备份迁移任务恢复目标RDS for SQL Server实例ID。
	TargetInstanceId string `json:"target_instance_id"`

	// 目标实例是否开启FileStream模式。可通过RDS for SQL Server详情接口获取。
	MsFileStreamStatus *string `json:"ms_file_stream_status,omitempty"`

	// RDS for SQL Server备份文件的文件ID，通过RDS全量恢复时必填。可通过云数据库RDS备份管理页面获取。
	FileId *string `json:"file_id,omitempty"`
}

func (o BackupJobEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupJobEndpointInfo struct{}"
	}

	return strings.Join([]string{"BackupJobEndpointInfo", string(data)}, " ")
}
