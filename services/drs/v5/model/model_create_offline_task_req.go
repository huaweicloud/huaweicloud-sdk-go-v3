package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOfflineTaskReq 创建备份迁移任务请求。
type CreateOfflineTaskReq struct {
	BaseInfo *BackupJobBaseInfo `json:"base_info"`

	TargetDbInfo *BackupJobEndpointInfo `json:"target_db_info"`

	BackupInfo *BackupInfo `json:"backup_info"`

	Options *BackupRestoreOptionInfo `json:"options"`
}

func (o CreateOfflineTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOfflineTaskReq struct{}"
	}

	return strings.Join([]string{"CreateOfflineTaskReq", string(data)}, " ")
}
