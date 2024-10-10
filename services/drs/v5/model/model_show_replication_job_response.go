package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReplicationJobResponse Response Object
type ShowReplicationJobResponse struct {

	// 任务ID。
	Id *string `json:"id,omitempty"`

	// 任务状态。
	Status *string `json:"status,omitempty"`

	// 任务创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 任务完成时间。
	FinishTime *string `json:"finish_time,omitempty"`

	BackupInfo *BackupInfoResp `json:"backup_info,omitempty"`

	BaseInfo *BackupJobBaseInfo `json:"base_info,omitempty"`

	TargetDbInfo *BackupJobEndpointInfo `json:"target_db_info,omitempty"`

	Options *BackupRestoreOptionInfo `json:"options,omitempty"`

	// 备份恢复数据库映射新名称。
	NewDbNames *string `json:"new_db_names,omitempty"`

	// RDS实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 迁移过程中失败原因。
	ErrorLog       *string `json:"error_log,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowReplicationJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationJobResponse struct{}"
	}

	return strings.Join([]string{"ShowReplicationJobResponse", string(data)}, " ")
}
