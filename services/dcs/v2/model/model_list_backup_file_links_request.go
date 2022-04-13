package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBackupFileLinksRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 备份记录ID。

	BackupId string `json:"backup_id"`

	Body *DownloadBackupFilesReq `json:"body,omitempty"`
}

func (o ListBackupFileLinksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupFileLinksRequest struct{}"
	}

	return strings.Join([]string{"ListBackupFileLinksRequest", string(data)}, " ")
}
