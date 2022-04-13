package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteManualBackupRequest struct {
	// 备份文件ID。

	BackupId string `json:"backup_id"`
}

func (o DeleteManualBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteManualBackupRequest struct{}"
	}

	return strings.Join([]string{"DeleteManualBackupRequest", string(data)}, " ")
}
