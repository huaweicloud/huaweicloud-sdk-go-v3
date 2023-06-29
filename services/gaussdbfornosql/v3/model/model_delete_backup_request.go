package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackupRequest Request Object
type DeleteBackupRequest struct {

	// 备份文件ID。
	BackupId string `json:"backup_id"`
}

func (o DeleteBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackupRequest", string(data)}, " ")
}
